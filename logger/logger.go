package logger

import (
	"fmt"
	"os"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Log *zap.Logger
	logFile *os.File
	fileMutex sync.Mutex
	maxSize = 3 * 1024 
	currentLog string
)

func InitLogger() {
	// Create logs directory if it doesn't exist
	if _, err := os.Stat("logger/logs"); os.IsNotExist(err) {
		os.MkdirAll("logger/logs", os.ModePerm)
	}

	// Create a new log file with rotation
	createNewLogFile()

	// Configure zap logger
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

	// Configure json logger with file rotation
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(os.Stdout),
			zapcore.AddSync(logFile),
		),
		zap.InfoLevel,
	)

	Log = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

	// Start log rotation monitoring
	go monitorLogSize()
}

func createNewLogFile() {
	fileMutex.Lock()
	defer fileMutex.Unlock()

	if logFile != nil {
		logFile.Close()
	}

	newName := time.Now().Format("2006-01-02_15-04-05")
	currentLog = fmt.Sprintf("logger/logs/%s.log", newName)
	var err error
	logFile, err = os.Create(currentLog)
	if err != nil {
		panic(err)
	}
}

func monitorLogSize() {
	for {
		time.Sleep(1 * time.Second)
		fileMutex.Lock()
		info, err := os.Stat(currentLog)
		if err == nil && info.Size() >= int64(maxSize) {
			createNewLogFile()
		}
		fileMutex.Unlock()
	}
}

func Sync() {
	fileMutex.Lock()
	defer fileMutex.Unlock()
	if Log != nil {
		Log.Sync()
	}
	if logFile != nil {
		logFile.Close()
	}
}