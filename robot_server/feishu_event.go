package robotServer

import "github.com/sirupsen/logrus"

type CallbackType func(map[string]any)

var eventMap = make(map[string]CallbackType)

func registerListener(f CallbackType, eventType string) {
	// Register a handler for a specific Feishu event

	// usage:

	// example handler:
	// func OnEventA(data map[string]any) {}

	// add this line to init():
	// RegisterListener(OnEventA, eventType)

	if _, isEventExist := eventMap[eventType]; isEventExist {
		logrus.Warning("Double declaration of event listener: ", eventType)
	}
	eventMap[eventType] = f
}
