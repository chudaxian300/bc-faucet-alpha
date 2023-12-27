package utils

import (
	"api/conf"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"time"
)

var log *logrus.Logger

var logToFile *logrus.Logger

var loggerFile string

func SetLogFile(file string) {
	loggerFile = file
}

func init() {
	SetLogFile(filepath.Join(conf.Conf.MyLog.Path, conf.Conf.MyLog.Name))
}

func Log() *logrus.Logger {
	if conf.Conf.MyLog.Model == "file" {
		return logFile()
	} else {
		if log == nil {
			log = logrus.New()
			log.Out = os.Stdout
			log.Formatter = &logrus.JSONFormatter{
				TimestampFormat: "2006-01-02 15:04:05",
			}
			log.SetLevel(logrus.DebugLevel)
		}
	}
	return log
}

func logFile() *logrus.Logger {
	if logToFile == nil {
		logToFile = logrus.New()
		logToFile.SetLevel(logrus.DebugLevel)

		logWriter, _ := rotatelogs.New(
			loggerFile+"_%Y%m%d.log",
			rotatelogs.WithMaxAge(30*24*time.Hour),
			rotatelogs.WithRotationTime(24*time.Hour),
		)

		writeMap := lfshook.WriterMap{
			logrus.InfoLevel:  logWriter,
			logrus.FatalLevel: logWriter,
			logrus.DebugLevel: logWriter,
			logrus.WarnLevel:  logWriter,
			logrus.ErrorLevel: logWriter,
			logrus.PanicLevel: logWriter,
		}

		lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})

		logToFile.AddHook(lfHook)
	}
	return logToFile
}
