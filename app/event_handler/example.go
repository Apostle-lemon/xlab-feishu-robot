package eventHandler

import "github.com/sirupsen/logrus"

func example(event map[string]any) {
	logrus.WithFields(logrus.Fields{"event": event}).Info("Example event handler triggered")
}
