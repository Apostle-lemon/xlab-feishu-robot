package config

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

func SetupLogrus() {
	// log filename: log_yyyymmdd_hhmmss.{fractional seconds}
	logPath := "./log/log_" + time.Now().Format("20060102_150405.00000")
	logFile, err := os.OpenFile(logPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Failed to create log file: ", logPath)
	}
	logrus.SetOutput(io.MultiWriter(os.Stdout, logFile))
}
