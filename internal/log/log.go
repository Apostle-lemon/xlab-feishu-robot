package log

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
)

func SetupLogrus() {
	logPath := "./log/log_" + time.Now().Format("20060102_150405.00000")
	if err := os.MkdirAll(filepath.Dir(logPath), os.ModePerm); err != nil {
		logrus.Error("Failed to create log directory: ", filepath.Dir(logPath))
		logrus.SetOutput(io.MultiWriter(os.Stdout))
		return
	} else {
		logFile, err := os.OpenFile(logPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
		if err != nil {
			fmt.Println("Failed to create log file: ", logPath)
			logrus.SetOutput(io.MultiWriter(os.Stdout))
		} else {
			logrus.SetOutput(io.MultiWriter(os.Stdout, logFile))
		}
		return
	}
}
