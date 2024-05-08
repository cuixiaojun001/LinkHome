// Package zapLog rotate function based on  lumberjack
package log

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// ensure we always implement io.WriteCloser
var _ io.WriteCloser = (*LogRotate)(nil)

// LogRotate is an io.WriteCloser that writes to the specified filename.
type LogRotate struct {
	// Filename is the file to write logs to.  Backup log files will be retained
	// in the same directory.  It uses <processname>-lumberjack.log in
	// os.TempDir() if empty.
	Filename string `json:"filename" yaml:"filename"`

	// Rotate policy.default based on size
	PolicyType TargetPolicyType `json:"rotate_type" yaml:"rotate_type"`

	MaxFileSize int `json:"max_file_size" yaml:"max_file_size"`

	// LocalTime determines if the time used for formatting the timestamps in
	// backup files is the computer's local time.  The default is to use UTC
	// time.
	LocalTime bool `json:"localtime" yaml:"localtime"`

	TimeProvider TimeProvider

	nextRotateAt    time.Time
	currentFilename string
	size            int64
	file            *os.File
	mu              sync.Mutex
}

type TimeProvider interface {
	Now() time.Time
}

var (
	// osStat exists so it can be mocked out by tests.
	osStat = os.Stat

	// megabyte is the conversion factor between MaxSize and bytes.  It is a
	// variable so tests can mock it out and not need to write megabytes of data
	// to disk.
	megabyte = 1024 * 1024
)

// Write implements io.Writer.  If a write would cause the log file to be larger
// than MaxSize, the file is closed, renamed to include a timestamp of the
// current time, and a new log file is created using the original log file name.
// If the length of the write is greater than MaxSize, an error is returned.
func (l *LogRotate) Write(p []byte) (n int, err error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.file == nil {
		if err = l.openNew(); err != nil {
			return 0, err
		}
	}

	if l.PolicyType == FileTargetSizeRotate {
		writeLen := int64(len(p))
		if writeLen > l.max() {
			return 0, fmt.Errorf(
				"write length %d exceeds maximum file size %d", writeLen, l.max(),
			)
		}

		if l.size+writeLen > l.max() {
			if err := l.rotate(); err != nil {
				return 0, err
			}
		}
	} else {
		if l.TimeProvider.Now().After(l.nextRotateAt) {
			if err := l.rotate(); err != nil {
				return 0, err
			}

		}
	}
	n, err = l.file.Write(p)
	l.size += int64(n)

	return n, err
}

// Close implements io.Closer, and closes the current logfile.
func (l *LogRotate) Close() error {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.close()
}

// close closes the file if it is open.
func (l *LogRotate) close() error {
	if l.file == nil {
		return nil
	}
	err := l.file.Close()
	l.file = nil
	return err
}

// Rotate causes Logger to close the existing log file and immediately create a
// new one.  This is a helper function for applications that want to initiate
// rotations outside of the normal rotation rules, such as in response to
// SIGHUP.  After rotating, this initiates compression and removal of old log
// files according to the configuration.
func (l *LogRotate) Rotate() error {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.rotate()
}

// rotate closes the current file, moves it aside with a timestamp in the name,
// (if it exists), opens a new file with the original filename, and then runs
// post-rotation processing and removal.
func (l *LogRotate) rotate() error {
	if err := l.close(); err != nil {
		return err
	}
	if err := l.openNew(); err != nil {
		return err
	}
	return nil
}

// openNew opens a new log file for writing, moving any old log file out of the
// way.  This methods assumes the file has already been closed.
func (l *LogRotate) openNew() error {
	err := os.MkdirAll(l.dir(), 0744)
	if err != nil {
		return fmt.Errorf("can't make directories for new logfile: %s", err)
	}

	name := l.filename()
	mode := os.FileMode(0644)

	f, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY|os.O_APPEND, mode)
	if err != nil {
		return fmt.Errorf("can't open new logfile: %s", err)
	}
	l.file = f
	l.size = 0
	return nil
}

// backupName creates a new filename from the given name, inserting a timestamp
// between the filename and the extension, using the local time if requested
// (otherwise UTC).
func (l *LogRotate) backupName(name string, local bool) string {
	dir := filepath.Dir(name)
	filename := filepath.Base(name)
	ext := filepath.Ext(filename)
	prefix := filename[:len(filename)-len(ext)]
	t := l.TimeProvider.Now()
	if !local {
		t = t.UTC()
	}

	var timeFormat string
	switch l.PolicyType {
	case FileTargetSizeRotate:
		timeFormat = sizeRotateTimeFormat
	case FileTargetHourRotate:
		timeFormat = hourRotateTimeFormat
	case FileTargetDayRotate:
		timeFormat = dayRotateTimeFormat
	}
	timestamp := t.Format(timeFormat)
	return filepath.Join(dir, fmt.Sprintf("%s-%s%s", prefix, timestamp, ext))
}

// genFilename generates the name of the logfile from the current time.
func (l *LogRotate) filename() string {
	if l.PolicyType != FileTargetSizeRotate {
		l.nextRotateAt = l.nextRotateTime()
	}
	if l.Filename != "" {
		l.currentFilename = l.backupName(l.Filename, l.LocalTime)
		return l.currentFilename
	}
	l.currentFilename = l.backupName("/tmp/ke.log", l.LocalTime)
	return l.currentFilename
}

func (l *LogRotate) nextRotateTime() time.Time {
	var future time.Time
	if l.PolicyType == FileTargetDayRotate {
		future = l.TimeProvider.Now().Add(time.Hour * 24)
		return time.Date(future.Year(), future.Month(), future.Day(), 0, 0, 0, 0, time.Local)
	} else {
		future = l.TimeProvider.Now().Add(time.Hour)
	}
	return time.Date(future.Year(), future.Month(), future.Day(), future.Hour(), 0, 0, 0, future.Location())
}

// max returns the maximum size in bytes of log files before rolling.
func (l *LogRotate) max() int64 {
	if l.MaxFileSize <= 0 {
		return int64(defaultFileTargetMaxSize * megabyte)
	}
	return int64(l.MaxFileSize) * int64(megabyte)
}

// dir returns the directory for the current filename.
func (l *LogRotate) dir() string {
	return filepath.Dir(l.Filename)
}
