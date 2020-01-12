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

	hook := lumberjack.Logger{
		Filename:   opts.FileName,
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
		EncodeTime:     jsonTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zapcore.Level(opts.Level))

	runMode := opts.RunMode
	var core zapcore.Core
	if runMode == constants.ReleaseMode {
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

	if opts.Config.RunMode == constants.DebugMode {
		caller := zap.AddCaller()
		development := zap.Development()
		logger = zap.New(core, caller, development)
	} else {
		logger = zap.New(core)
	}

	return logger, err
}

// jsonTimeEncoder 自定义时间格式化
func jsonTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006/01/05 15:04:05:000"))
}
