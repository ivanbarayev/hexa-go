package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"main/config"
	"os"
)

// Logger methods interface
type Logger interface {
	InitLogger()
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	DPanic(args ...interface{})
	DPanicf(template string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(template string, args ...interface{})
}

// Logger
type ApiLogger struct {
	cfg         *config.Config
	sugarLogger *zap.SugaredLogger
}

// App Logger constructor
func NewApiLogger(cfg *config.Config) *ApiLogger {
	return &ApiLogger{cfg: cfg}
}

// For mapping config logger to app logger levels
var loggerLevelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func (l *ApiLogger) getLoggerLevel(cfg *config.Config) zapcore.Level {
	level, exist := loggerLevelMap[cfg.Logger.LEVEL]
	if !exist {
		return zapcore.DebugLevel
	}

	return level
}

// Init logger
func (l *ApiLogger) InitLogger() {
	logLevel := l.getLoggerLevel(l.cfg)

	logWriter := zapcore.AddSync(os.Stderr)

	var encoderCfg zapcore.EncoderConfig
	if l.cfg.Server.APP_ENV == "Development" {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderCfg = zap.NewProductionEncoderConfig()
	}

	var encoder zapcore.Encoder
	encoderCfg.LevelKey = "LEVEL"
	encoderCfg.CallerKey = "CALLER"
	encoderCfg.TimeKey = "TIME"
	encoderCfg.NameKey = "NAME"
	encoderCfg.MessageKey = "MESSAGE"

	if l.cfg.Logger.ENCODING == "console" {
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	}

	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(encoder, logWriter, zap.NewAtomicLevelAt(logLevel))
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	l.sugarLogger = logger.Sugar()
	if err := l.sugarLogger.Sync(); err != nil {
		l.sugarLogger.Error(err)
	}
}

// Logger methods

func (l *ApiLogger) Debug(args ...interface{}) {
	l.sugarLogger.Debug(args...)
}

func (l *ApiLogger) Debugf(template string, args ...interface{}) {
	l.sugarLogger.Debugf(template, args...)
}

func (l *ApiLogger) Info(args ...interface{}) {
	l.sugarLogger.Info(args...)
}

func (l *ApiLogger) Infof(template string, args ...interface{}) {
	l.sugarLogger.Infof(template, args...)
}

func (l *ApiLogger) Warn(args ...interface{}) {
	l.sugarLogger.Warn(args...)
}

func (l *ApiLogger) Warnf(template string, args ...interface{}) {
	l.sugarLogger.Warnf(template, args...)
}

func (l *ApiLogger) Error(args ...interface{}) {
	l.sugarLogger.Error(args...)
}

func (l *ApiLogger) Errorf(template string, args ...interface{}) {
	l.sugarLogger.Errorf(template, args...)
}

func (l *ApiLogger) DPanic(args ...interface{}) {
	l.sugarLogger.DPanic(args...)
}

func (l *ApiLogger) DPanicf(template string, args ...interface{}) {
	l.sugarLogger.DPanicf(template, args...)
}

func (l *ApiLogger) Panic(args ...interface{}) {
	l.sugarLogger.Panic(args...)
}

func (l *ApiLogger) Panicf(template string, args ...interface{}) {
	l.sugarLogger.Panicf(template, args...)
}

func (l *ApiLogger) Fatal(args ...interface{}) {
	l.sugarLogger.Fatal(args...)
}

func (l *ApiLogger) Fatalf(template string, args ...interface{}) {
	l.sugarLogger.Fatalf(template, args...)
}
