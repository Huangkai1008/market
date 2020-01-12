package logging

import (
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"market/internal/pkg/config"
	"market/internal/pkg/constants"
)

// Options 日志配置可选项
type Options struct {
	*config.Config
}

// NewWithOptions 返回一个Zap Logger实例
func NewWithOptions(opts *Options) (*zap.Logger, error) {
	var (
		err    error
		logger *zap.Logger
	)

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     jsonTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zapcore.Level(opts.Level))

	cores := make([]zapcore.Core, 0, 2)

	je := zapcore.NewJSONEncoder(encoderConfig)
	hook := lumberjack.Logger{
		Filename:   opts.FileName,
		MaxSize:    128,
		MaxBackups: 30,
		MaxAge:     30,
		Compress:   true,
	}
	fileCore := zapcore.NewCore(je, zapcore.AddSync(&hook), atomicLevel)
	cores = append(cores, fileCore)

	var options []zap.Option
	if opts.Config.RunMode == constants.DebugMode {
		ce := zapcore.NewConsoleEncoder(encoderConfig)
		consoleCore := zapcore.NewCore(ce, zapcore.AddSync(os.Stdout), atomicLevel)
		cores = append(cores, consoleCore)
		caller := zap.AddCaller()
		development := zap.Development()
		options = append(options, caller, development)
	}

	core := zapcore.NewTee(cores...)
	logger = zap.New(core, options...)

	zap.ReplaceGlobals(logger)

	return logger, err
}

// jsonTimeEncoder 自定义时间格式化
func jsonTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006/01/05 15:04:05:000"))
}
