package logger

import (
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

type ActionENUM string
type LogTypeENUM string

const (
	Parse    ActionENUM = "parse"
	Validate ActionENUM = "validation"
	Usecase  ActionENUM = "usecase"

	Dev  LogTypeENUM = "dev"
	Prod LogTypeENUM = "prod"
)

type Data struct {
	Module string                 `json:"module"`
	Method string                 `json:"method"`
	Action ActionENUM             `json:"action"`
	Params map[string]interface{} `json:"params"`
}

// Interface -.
type Interface interface {
	Info(message string, args Data)
	Warn(message string, args Data)
	Error(message interface{}, args Data)
}

// Logger -.
type Logger struct {
	logger  *logrus.Logger
	typeLog LogTypeENUM
}

var _ Interface = (*Logger)(nil)

// New -.
func New() *Logger {
	var l logrus.Level
	logger := logrus.New()

	level, ok := os.LookupEnv("LOG_LEVEL")
	if !ok || len(level) == 0 {
		level = "dev"
	}

	switch level {
	case "prod":
		l = logrus.InfoLevel
	case "dev":
		l = logrus.DebugLevel
	default:
		l = logrus.DebugLevel
	}

	if l == logrus.InfoLevel {
		logFilePath := "logs/app.log"

		logFile := &lumberjack.Logger{
			Filename:   logFilePath,
			MaxSize:    10, // Максимальный размер файла в мегабайтах перед ротацией
			MaxBackups: 5,  // Максимальное количество резервных файлов
			MaxAge:     28, // Максимальный возраст файла в днях
			Compress:   true,
		}

		logger.SetLevel(l)

		logger.AddHook(lfshook.NewHook(
			lfshook.WriterMap{
				logrus.InfoLevel:  logFile,
				logrus.WarnLevel:  logFile,
				logrus.ErrorLevel: logFile,
			},
			&logrus.JSONFormatter{},
		))

		defer func(logFile *lumberjack.Logger) {
			err := logFile.Close()
			if err != nil {
				logrus.Error(err)
			}
		}(logFile)

	} else {
		logger.SetLevel(l)
	}

	return &Logger{
		logger: logger,
	}
}

// Info -.
func (l *Logger) Info(message string, args Data) {
	if l.logger.Level == logrus.DebugLevel {
		timeStamp := time.Now().Format("2006-01-02 15:04:05")
		l.logger.Info("INFO: ", timeStamp, " - ", message, args)
	}
}

// Warn -.
func (l *Logger) Warn(message string, args Data) {
	timeStamp := time.Now().Format("2006-01-02 15:04:05")
	l.logger.Warn("WARN: ", timeStamp, " - ", message, args)
}

// Error -.
func (l *Logger) Error(message interface{}, args Data) {
	timeStamp := time.Now().Format("2006-01-02 15:04:05")
	l.logger.Error("ERROR: ", timeStamp, " - ", message, args)
}
