package zap

import (
	"data_feedback_progress/public/config"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	FileRotateInfo   *lumberjack.Logger
	FileRotateError  *lumberjack.Logger
	AtomicLevelInfo  zap.AtomicLevel
	AtomicLevelError zap.AtomicLevel
)

var (
	zapLogger *zap.SugaredLogger
	zapOnce   sync.Once
)

func init() {

	FileRotateInfo = &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/runtime/logs/default.log", config.App().BasePath),
		MaxSize:    1000,
		MaxBackups: 7,
	}
	FileRotateError = &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/runtime/logs/default.error.log", config.App().BasePath),
		MaxSize:    1000,
		MaxBackups: 7,
	}

	AtomicLevelInfo = zap.NewAtomicLevelAt(zap.InfoLevel)
	AtomicLevelError = zap.NewAtomicLevelAt(zap.ErrorLevel)
}

func Zap() (logger *zap.SugaredLogger) {
	if zapLogger == nil {
		zapOnce.Do(func() {
			initZap()
		})
	}

	return zapLogger
}

func initZap() {
	WriterSyncersInfo := []zapcore.WriteSyncer{
		zapcore.AddSync(FileRotateInfo),
	}
	WriterSyncersError := []zapcore.WriteSyncer{
		zapcore.AddSync(FileRotateError),
	}

	if config.Debug {
		AtomicLevelInfo.SetLevel(zap.DebugLevel)
		WriterSyncersInfo = append(WriterSyncersInfo, zapcore.AddSync(os.Stdout))
		//WriterSyncersError = append(WriterSyncersError, zapcore.AddSync(os.Stdout)) // error不打控制台
	}

	coreInfo := zapcore.NewCore(
		zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
			TimeKey:       "T",
			LevelKey:      "L",
			NameKey:       "N",
			CallerKey:     "C",
			MessageKey:    "M",
			StacktraceKey: "S",
			LineEnding:    zapcore.DefaultLineEnding,
			EncodeLevel:   zapcore.CapitalLevelEncoder,
			EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
			},
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}),
		zapcore.NewMultiWriteSyncer(WriterSyncersInfo...),
		AtomicLevelInfo,
	)
	coreError := zapcore.NewCore(
		zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
			TimeKey:       "T",
			LevelKey:      "L",
			NameKey:       "N",
			CallerKey:     "C",
			MessageKey:    "M",
			StacktraceKey: "S",
			LineEnding:    zapcore.DefaultLineEnding,
			EncodeLevel:   zapcore.CapitalLevelEncoder,
			EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
			},
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}),
		zapcore.NewMultiWriteSyncer(WriterSyncersError...),
		AtomicLevelError,
	)
	core := zapcore.NewTee(coreInfo, coreError)
	logger := zap.New(core, zap.AddCaller())

	zapLogger = logger.Sugar()
}

type ZapOutput struct {
	Logger *zap.SugaredLogger
}

func (t *ZapOutput) Write(p []byte) (n int, err error) {
	t.Logger.Info(string(p))
	return len(p), nil
}
