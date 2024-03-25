package logger

import (
	"errors"
	"os"
	"path"

	"github.com/cuixiaojun001/linkhome/common/config"
	"github.com/cuixiaojun001/linkhome/common/env"
	klog "github.com/cuixiaojun001/linkhome/library/logger"
	"github.com/cuixiaojun001/linkhome/library/utils"
)

var (
	accessLogger *klog.RawLogger
	appLogger    klog.Logger
)

// Package Level Functions
var (
	Debugw func(string, ...interface{})
	Infow  func(string, ...interface{})
	Warnw  func(string, ...interface{})
	Errorw func(string, ...interface{})
	Panicw func(string, ...interface{})
	Fatalw func(string, ...interface{})
	With   func(...interface{}) klog.Logger
)

// LogConfig logger配置信息结构
type LogConfig struct {
	LogFile      string                // LogFile 日志文件
	LogLevel     klog.Level            // LogLevel 日志输出级别
	RotatePolicy klog.TargetPolicyType // RotatePolicy 日志Rotate策略
	MaxFileSize  int                   // MaxFileSize  最大日志文件大小
	IsShowCaller bool                  // IsShowCaller 是否显示Caller信息
	CallerSkip   int                   // CallerSkip Caller Skip，参见 runtime.Caller
}

func SetUp() {
	SetUpAccessLog("log.storage.access") // access访问日志
	SetUpAppLog("log.storage.app")       // app日志
}

func SetUpAccessLog(accessLogKey string) {
	if config.IsSet(accessLogKey) {
		accessConf, err := GetLogConfig(accessLogKey)
		if err != nil {
			panic(err)
		} else {
			target := klog.NewTarget(klog.FileTarget, accessConf.RotatePolicy, klog.FormatterRaw, klog.WithFilename(accessConf.LogFile), klog.WithMaxFileSize(accessConf.MaxFileSize))
			kLog := klog.NewZapLogger(accessConf.LogLevel, klog.AddTarget(target))
			accessLogger = klog.NewRawLogger(kLog)
		}
	}
}

func SetUpAppLog(appKey string) {
	logConf, err := GetLogConfig(appKey)
	if err != nil {
		panic(err)
	} else {
		switch env.Env {
		case env.Develop, env.Test: // 测试，开发环境打日志到Console和File
			consoleTarget := klog.NewTarget(klog.ConsoleTarget, logConf.RotatePolicy, klog.FormatterRaw)
			fileTarget := klog.NewTarget(klog.FileTarget, logConf.RotatePolicy, klog.FormatterJson, klog.WithFilename(logConf.LogFile), klog.WithMaxFileSize(logConf.MaxFileSize))
			logger := klog.NewZapLogger(logConf.LogLevel, klog.AddTarget(consoleTarget), klog.AddTarget(fileTarget), klog.SetCaller(logConf.IsShowCaller, logConf.CallerSkip))
			appLogger = logger
		case env.Product, env.Preview: // 生产，预发环境打日志到File
			target := klog.NewTarget(klog.FileTarget, logConf.RotatePolicy, klog.FormatterJson, klog.WithFilename(logConf.LogFile), klog.WithMaxFileSize(logConf.MaxFileSize))
			logger := klog.NewZapLogger(logConf.LogLevel, klog.AddTarget(target), klog.SetCaller(logConf.IsShowCaller, logConf.CallerSkip))
			appLogger = logger
		default:
			panic("env is not support")
		}
	}
	Debugw = appLogger.Debugw
	Infow = appLogger.Infow
	Warnw = appLogger.Warnw
	Errorw = appLogger.Errorw
	Panicw = appLogger.Panicw
	Fatalw = appLogger.Fatalw
	With = appLogger.With
}

// GetLogConfig 根据key获取日志配置
func GetLogConfig(key string) (*LogConfig, error) {
	var logConf = LogConfig{}
	if !config.IsSet(key) {
		return nil, errors.New("config key is not existed")
	}
	err := config.Sub(key).Unmarshal(&logConf)
	if err != nil {
		return nil, err
	}
	if logConf.LogFile == "" {
		return nil, errors.New("logfile is not define")
	}
	if logConf.LogLevel == "" {
		logConf.LogLevel = klog.Debug
	}
	if logConf.RotatePolicy == "" {
		logConf.RotatePolicy = klog.FileTargetHourRotate
	}
	logPath := path.Dir(logConf.LogFile)
	if !utils.IsDir(logPath) {
		err := os.MkdirAll(logPath, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	return &logConf, nil
}

// GetAccessLogger 获得Access Logger对象，需要先调用setUpxxx()初始化Logger
func GetAccessLogger() *klog.RawLogger {
	return accessLogger
}

// GetAppLogger 获得App Logger对象，需要先调用setUpxxx()初始化Logger
func GetAppLogger() klog.Logger {
	return appLogger
}
