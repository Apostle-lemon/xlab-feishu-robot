package exampleHandler

import "github.com/sirupsen/logrus"

func Handler(data map[string]any) {
	logrus.WithFields(logrus.Fields{"data": data}).Info("Example event handler triggered")
}
