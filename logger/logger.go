package logger

import (
	"fmt"
	"os"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type RotatingWriter struct {
	sync.Mutex
	dir         string
	maxSize     int64
	currentFile *os.File
	currentSize int64
	startTime   time.Time
}

func NewRotatingWriter(dir string, maxSize int64) *RotatingWriter {
	rw := &RotatingWriter{
		dir:     dir,
		maxSize: maxSize,
	}
	rw.createNewFile()
	return rw
}

func (rw *RotatingWriter) Write(p []byte) (n int, err error) {
	rw.Lock()
	defer rw.Unlock()

	n, err = rw.currentFile.Write(p)
	if err != nil {
		return n, err
	}

	rw.currentSize += int64(n)
	if rw.currentSize >= rw.maxSize {
		rw.createNewFile()
	}
	return n, err
}

func (rw *RotatingWriter) createNewFile() {
	if rw.currentFile != nil {
		rw.currentFile.Close()
	}

	rw.startTime = time.Now()
	name := fmt.Sprintf("%s/%s.log", rw.dir, rw.startTime.Format("2006-01-02_15-04-05"))
	file, err := os.Create(name)
	if err != nil {
		panic(err)
	}

	rw.currentFile = file
	rw.currentSize = 0
	fmt.Println("Created new log file: ", name)
}

func (rw *RotatingWriter) Close() {
	rw.Lock()
	defer rw.Unlock()

	if rw.currentFile != nil {
		rw.currentFile.Close()
	}
}

var Log *zap.Logger
var writer *RotatingWriter

func InitLogger() {
	// Create logs directory if it doesn't exist
	logDir := "logger/logs"
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		os.MkdirAll(logDir, os.ModePerm)
	}

	writer = NewRotatingWriter(logDir, 3*1024)

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
			zapcore.AddSync(writer),
		),
		zap.InfoLevel,
	)

	Log = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
}

func Sync() {
	if Log != nil {
		Log.Sync()
	}
	if writer != nil {
		writer.Close()
	}
}
