package conf

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
)

func setupLogger() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(loggerSetting.getLogLevel())

	// Log output to file
	out := newFileLogger()
	logrus.SetOutput(out)
}

func newFileLogger() io.Writer {
	filename := fmt.Sprintf("%s/%s%s",
		fileLoggerSetting.SavePath,
		fileLoggerSetting.FileName,
		fileLoggerSetting.FileExt)
	return &lumberjack.Logger{
		Filename:  filename,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}
}
