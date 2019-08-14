package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"market/pkg/setting"
	"os"
)

// gin middleware log
var GinLogger *zap.Logger

// SetupLogger setup logger
func init() {
	GinLogger = initLogger(setting.GinLogPath, zap.InfoLevel, false)
}

// InitLogger initial the zap logger config
func initLogger(logPath string, logLevel zapcore.Level, isDev bool) *zap.Logger {
	var logger *zap.Logger

	hook := lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    128,
		MaxBackups: 30,
		MaxAge:     30,
		Compress:   true,
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(logLevel)

	runMode := setting.RunMode
	var core zapcore.Core
	if runMode == "release" {
		core = zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook)),
			atomicLevel,
		)
	} else {
		core = zapcore.NewCore(
			zapcore.NewJSONEncoder(encoderConfig),
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)),
			atomicLevel,
		)
	}

	if isDev {
		// catch line num and stack info
		caller := zap.AddCaller()
		development := zap.Development()
		logger = zap.New(core, caller, development)
	} else {
		logger = zap.New(core)
	}

	return logger
}
