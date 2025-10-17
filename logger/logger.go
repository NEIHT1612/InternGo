package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Log *zap.Logger

func InitLogger() {
	if _, err := os.Stat("logger/logs"); os.IsNotExist(err) {
		os.MkdirAll("logger/logs", os.ModePerm)
	}

	logFile := &lumberjack.Logger{
		Filename:   "logger/logs/app.log", 
		MaxSize:    10,  // Megabytes
		MaxBackups: 5,  // Keep last 5 log files
		MaxAge:     30,  // Days to keep log files
		Compress:   true, // Compress old log files
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(os.Stdout),
			zapcore.AddSync(logFile),
		),
		zap.InfoLevel,
	)

	Log = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
}

func Sync() {
	Log.Sync()
}