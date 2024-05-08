package log

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ZapLogger implement Logger.delegate to zap,so we can use zap's function directly
type ZapLogger struct {
	caller      bool
	callerSkip  int
	logRotate   *LogRotate
	targets     []*Target
	logger      *zap.Logger
	sugerLogger *zap.SugaredLogger
}

// SetCaller set caller
func (zapLog *ZapLogger) SetCaller(c bool, callerSkip int) {
	zapLog.caller = c
	if callerSkip <= 0 {
		zapLog.callerSkip = 0
	} else {
		zapLog.callerSkip = callerSkip
	}
}

// GetCaller get caller
func (zapLog *ZapLogger) GetCaller() bool {
	return zapLog.caller
}

type Timer struct {
}

func (t *Timer) Now() time.Time {
	return time.Now()
}

// NewLogger make a Klogger
func NewZapLogger(level Level, opts ...LoggerOption) Logger {
	var encoder zapcore.Encoder
	zapLog := &ZapLogger{}
	for _, opt := range opts {
		opt.apply(zapLog)
	}
	cores := make([]zapcore.Core, 0)
	for _, t := range zapLog.targets {
		switch t.Type {
		case ConsoleTarget:
			encoderConfig := zap.NewProductionEncoderConfig()
			encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
			encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
			switch t.Formatter {
			case FormatterRaw:
				encoder = zapcore.NewConsoleEncoder(encoderConfig)
			case FormatterJson:
				encoder = zapcore.NewJSONEncoder(encoderConfig)
			}
			writeSyncer := zapcore.AddSync(os.Stdout)
			cores = append(cores, zapcore.NewCore(encoder, writeSyncer, zapLog.getLevel(level)))
		case FileTarget:
			logRotate := &LogRotate{
				Filename:     t.FileName,
				PolicyType:   t.TargetPolicyType,
				LocalTime:    true,
				TimeProvider: &Timer{},
			}
			encoderConfig := zap.NewProductionEncoderConfig()
			encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
			encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
			switch t.Formatter {
			case FormatterRaw:
				encoder = NewZapRawEncoder(encoderConfig)
			case FormatterJson:
				encoder = zapcore.NewJSONEncoder(encoderConfig)
			}
			switch t.TargetPolicyType {
			case FileTargetSizeRotate:
				logRotate.MaxFileSize = t.MaxFileSize
				writeSyncer := zapcore.AddSync(logRotate)
				cores = append(cores, zapcore.NewCore(encoder, writeSyncer, zapLog.getLevel(level)))
			case FileTargetHourRotate, FileTargetDayRotate:
				writeSyncer := zapcore.AddSync(logRotate)
				cores = append(cores, zapcore.NewCore(encoder, writeSyncer, zapLog.getLevel(level)))
			}
		default:
			panic("not support target type")
		}
	}
	core := zapcore.NewTee(cores...)
	var zlog *zap.Logger
	if zapLog.caller {
		zlog = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(zapLog.callerSkip), zap.AddStacktrace(zapcore.PanicLevel))
	} else {
		zlog = zap.New(core, zap.AddStacktrace(zapcore.PanicLevel))
	}
	zapLog.logger = zlog
	zapLog.sugerLogger = zlog.Sugar()
	return zapLog
}

//// NewZapLogger make a Klogger
//func NewZapLogger(level Level, opts ...LoggerOption) Logger {
//	var encoder zapcore.Encoder
//	zapLog := &ZapLogger{}
//	for _, opt := range opts {
//		opt.apply(zapLog)
//	}
//	cores := make([]zapcore.Core, 0)
//	for _, t := range zapLog.targets {
//		switch t.Formatter {
//		case FormatterRaw:
//			consoleCfg := zap.NewProductionEncoderConfig()
//			consoleCfg.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
//				enc.AppendString(t.Format("2006-01-02 15:04:05"))
//			}
//			consoleCfg.EncodeLevel = zapcore.CapitalLevelEncoder
//			consoleCfg.EncodeCaller = func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
//				enc.AppendString(fmt.Sprintf("%s:%d", filepath.Base(caller.FullPath()), caller.Line))
//			}
//			encoder = zapcore.NewConsoleEncoder(consoleCfg)
//		case FormatterJson:
//			fileCfg := zap.NewProductionEncoderConfig()
//			fileCfg.EncodeTime = zapcore.ISO8601TimeEncoder
//			fileCfg.EncodeLevel = zapcore.CapitalLevelEncoder
//			encoder = zapcore.NewJSONEncoder(fileCfg)
//		}
//		switch t.Type {
//		case ConsoleTarget:
//			// writeSyncer := zapcore.AddSync(os.Stdout)
//			cores = append(cores, zapcore.NewCore(
//				encoder,
//				zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)),
//				zapLog.getLevel(level)),
//			)
//		case FileTarget:
//			logRotate := &LogRotate{
//				Filename:     t.FileName,
//				PolicyType:   t.TargetPolicyType,
//				LocalTime:    true,
//				TimeProvider: &Timer{},
//			}
//			switch t.TargetPolicyType {
//			case FileTargetSizeRotate:
//				logRotate.MaxFileSize = t.MaxFileSize
//				writeSyncer := zapcore.AddSync(logRotate)
//				cores = append(cores, zapcore.NewCore(encoder, writeSyncer, zapLog.getLevel(level)))
//			case FileTargetHourRotate, FileTargetDayRotate:
//				writeSyncer := zapcore.AddSync(logRotate)
//				cores = append(cores, zapcore.NewCore(encoder, writeSyncer, zapLog.getLevel(level)))
//			}
//		}
//	}
//	core := zapcore.NewTee(cores...)
//	var zlog *zap.Logger
//	if zapLog.caller {
//		zlog = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(zapLog.callerSkip), zap.AddStacktrace(zapcore.PanicLevel))
//	} else {
//		zlog = zap.New(core, zap.AddStacktrace(zapcore.PanicLevel))
//	}
//	zapLog.logger = zlog
//	zapLog.sugerLogger = zlog.Sugar()
//	return zapLog
//}

func (zapLog *ZapLogger) AddTarget(t *Target) {
	zapLog.targets = append(zapLog.targets, t)
}

func (zapLog *ZapLogger) getLevel(level Level) zapcore.Level {
	switch level {
	case Debug:
		return zapcore.DebugLevel
	case Info:
		return zapcore.InfoLevel
	case Warn:
		return zapcore.WarnLevel
	case Error:
		return zapcore.ErrorLevel
	case Panic:
		return zapcore.PanicLevel
	case Fatal:
		return zapcore.FatalLevel
	}
	return zapcore.InfoLevel
}

// Debugw use zap's sugar Debugw
func (zapLog *ZapLogger) Debugw(msg string, keysAndValues ...interface{}) {
	if zapLog == nil || zapLog.sugerLogger == nil {
		return
	}
	zapLog.sugerLogger.Debugw(msg, keysAndValues...)
	return
}

// Infow use zap's sugar Infow
func (zapLog *ZapLogger) Infow(msg string, keysAndValues ...interface{}) {
	if zapLog == nil || zapLog.sugerLogger == nil {
		return
	}
	zapLog.sugerLogger.Infow(msg, keysAndValues...)
	return
}

// Warnw use zap's sugar Warnw
func (zapLog *ZapLogger) Warnw(msg string, keysAndValues ...interface{}) {
	if zapLog == nil || zapLog.sugerLogger == nil {
		return
	}
	zapLog.sugerLogger.Warnw(msg, keysAndValues...)
	return
}

// Errorw use zap's sugar Errorw
func (zapLog *ZapLogger) Errorw(msg string, keysAndValues ...interface{}) {
	if zapLog == nil || zapLog.sugerLogger == nil {
		return
	}
	zapLog.sugerLogger.Errorw(msg, keysAndValues...)
	return
}

// Panicw use zap's sugar Panicw
func (zapLog *ZapLogger) Panicw(msg string, keysAndValues ...interface{}) {
	zapLog.sugerLogger.Panicw(msg, keysAndValues...)
	return
}

// Fatalw use zap's sugar Fatalw
func (zapLog *ZapLogger) Fatalw(msg string, keysAndValues ...interface{}) {
	if zapLog == nil || zapLog.sugerLogger == nil {
		return
	}
	zapLog.sugerLogger.Fatalw(msg, keysAndValues...)
	return
}

// With add additional fields
func (zapLog *ZapLogger) With(args ...interface{}) Logger {
	if zapLog.sugerLogger == nil {
		return zapLog
	}
	zapLog.sugerLogger = zapLog.sugerLogger.With(args...)
	return zapLog
}
