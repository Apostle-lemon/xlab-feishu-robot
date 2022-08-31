package dispatcher

import "github.com/sirupsen/logrus"

func RegisterListener(f CallbackType, eventType string) {
	// Register a handler for a specific Feishu event

	if _, isEventExist := eventMap[eventType]; isEventExist {
		logrus.Warning("Double declaration of event listener: ", eventType)
	}
	eventMap[eventType] = f
}
