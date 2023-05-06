package logging

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

type SimpleFormatter struct{}

func (f *SimpleFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b bytes.Buffer
	fmt.Fprintf(&b, "%s %s\n", entry.Time.Format(time.RFC3339), entry.Message)
	return b.Bytes(), nil
}

func NewLogger() (*logrus.Logger, error) {
	logFile := &lumberjack.Logger{
		Filename:   "deliverhalf.log",
		MaxSize:    10, // In megabytes
		MaxBackups: 0,
		MaxAge:     365, // In days
		Compress:   true,
	}

	logger := logrus.New()
	logger.SetOutput(io.MultiWriter(os.Stdout, logFile))
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: time.RFC3339,
		DisableColors:   true,
		FullTimestamp:   true,
	})
	// logger.SetFormatter(&logrus.JSONFormatter{})
	// logger.SetFormatter(&logrus.TextFormatter{})

	// formatter := &logrus.TextFormatter{
	//     TimestampFormat: "2006-01-02 15:04:05",
	// }
	logger.SetFormatter(&SimpleFormatter{})
	return logger, nil
}

var Logger, _ = NewLogger()
