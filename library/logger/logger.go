package log

import (
	"fmt"
	"strings"
	"sync"
)

// Formatter describes the formatter of a log message.
type Formatter int

const (
	FormatterJson Formatter = iota
	FormatterRaw
)

var FormatterNames = map[Formatter]string{
	FormatterJson: "Json",
	FormatterRaw:  "Raw",
}

// String returns the string representation of the log level
func (w Formatter) String() string {
	if name, ok := FormatterNames[w]; ok {
		return name
	}
	return FormatterNames[FormatterJson]
}

func StringToFormatter(name string) Formatter {
	for formatter, formatterName := range FormatterNames {
		if strings.ToLower(name) == strings.ToLower(formatterName) {
			return formatter
		}
	}
	return FormatterJson
}

// TargetPolicyType rotate policy
type TargetPolicyType string

const (
	// FileTargetSizeRotate Rotate log based on size,default size 100M
	FileTargetSizeRotate TargetPolicyType = "FILE_TARGET_SIZE"
	// FileTargetHourRotate Rotate log based on hour
	FileTargetHourRotate TargetPolicyType = "FILE_TARGET_HOUR"
	// FileTargetDayRotate Rotate log based on day
	FileTargetDayRotate TargetPolicyType = "FILE_TARGET_DAY"
)

// Level log level
type Level string

const (
	Debug Level = "debug"
	Info  Level = "info"
	Warn  Level = "warn"
	Error Level = "error"
	Panic Level = "panic"
	Fatal Level = "fatal"
)

type Target struct {
	Type             TargetType
	TargetPolicyType TargetPolicyType
	FileName         string
	MaxFileSize      int
	Formatter        Formatter
}

type TargetType string

const (
	FileTarget    TargetType = "file"
	ConsoleTarget TargetType = "console"
	EmailTarget   TargetType = "email"
	KafkaTarget   TargetType = "kafka"
)

const (
	sizeRotateTimeFormat          = "20060102150405.000"
	hourRotateTimeFormat          = "2006010215"
	dayRotateTimeFormat           = "20060102"
	defaultFileTargetRotatePolicy = FileTargetHourRotate
	defaultFileTargetMaxSize      = 100
	defaultCaller                 = true
)

var loggers = make(map[string]Logger)
var rwlock = new(sync.RWMutex)

// Logger implemented by base log package,like zap,logrus ...
type Logger interface {
	Debugw(msg string, keysAndValues ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	Panicw(msg string, keysAndValues ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})
	With(args ...interface{}) Logger

	AddTarget(target *Target)
	SetCaller(caller bool, callerSkip int)
}

type RawLogger struct {
	log Logger
}

// LoggerOption set option
type LoggerOption interface {
	apply(Logger)
}

// OptionFunc implement LoggerOption
type OptionFunc func(Logger)

func (f OptionFunc) apply(log Logger) {
	f(log)
}

func AddTarget(t *Target) LoggerOption {
	return OptionFunc(func(log Logger) {
		log.AddTarget(t)
	})
}
func SetCaller(caller bool, callerSkip int) LoggerOption {
	return OptionFunc(func(log Logger) {
		log.SetCaller(caller, callerSkip)
	})
}

// TargetOption set option
type TargetOption interface {
	apply(*Target)
}

// TargetOptionFunc implement LoggerOption
type TargetOptionFunc func(*Target)

func (f TargetOptionFunc) apply(target *Target) {
	f(target)
}

func WithFilename(filename string) TargetOption {
	return TargetOptionFunc(func(t *Target) {
		t.FileName = filename
	})
}

func WithMaxFileSize(maxFileSize int) TargetOption {
	return TargetOptionFunc(func(t *Target) {
		t.MaxFileSize = maxFileSize
	})
}

func NewTarget(typ TargetType, policy TargetPolicyType, formatter Formatter, opts ...TargetOption) *Target {
	t := &Target{
		Type:             typ,
		TargetPolicyType: policy,
		Formatter:        formatter,
	}
	for _, opt := range opts {
		opt.apply(t)
	}
	return t
}

func NewLogger(name string, log Logger, target *Target) Logger {
	logger := GetLogger(name)
	if logger != nil {
		return logger
	}
	rwlock.Lock()
	loggers[name] = log
	logger.AddTarget(target)
	rwlock.Unlock()
	return logger
}

// Logger get a logger by name.if name doesn't exist,then return nil.so be careful of panic
func GetLogger(name string) Logger {
	rwlock.RLock()
	logger, ok := loggers[name]
	if !ok {
		rwlock.RUnlock()
		return nil
	}
	rwlock.RUnlock()
	return logger
}

func NewRawLogger(logger Logger) *RawLogger {
	return &RawLogger{log: logger}
}

// Log msg
func (logger *RawLogger) Logf(format string, args ...interface{}) {
	if logger == nil {
		return
	}
	msg := fmt.Sprintf(format, args...)
	logger.log.Infow(msg)
}
