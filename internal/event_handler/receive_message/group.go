package receiveMessage

import (
	"github.com/sirupsen/logrus"
	"strings"
)

var groupMessageMap = make(map[string]messageHandler)

func group(messageevent *MessageEvent){
	switch strings.ToUpper(messageevent.Message.Message_type) {
	case "TEXT":
		groupTextMessage(messageevent)
	default:
		logrus.WithFields(logrus.Fields{"message type": messageevent.Message.Message_type}).Warn("Receive group message, but this type is not supported")
	}
}

func groupTextMessage(messageevent *MessageEvent){
	// get the pure text message, without @xxx
	messageevent.Message.Content = strings.TrimSuffix(strings.TrimPrefix(messageevent.Message.Content, "{\"text\":\""), "\"}")
	messageevent.Message.Content = messageevent.Message.Content[strings.Index(messageevent.Message.Content, " ")+1:]
	logrus.WithFields(logrus.Fields{"message content": messageevent.Message.Content}).Info("Receive group TEXT message")

	if handler, exists := groupMessageMap[messageevent.Message.Content]; exists {
		handler(messageevent)
		return
	} else {
		logrus.Error("Group message failed to find event handler: ", messageevent.Message.Content)
		return
	}
}

func groupMessageRegister(f messageHandler, s string) {
	
	if _, isEventExist := groupMessageMap[s]; isEventExist {
		logrus.Warning("Double declaration of group message handler: ", s)
	}
	groupMessageMap[s] = f
}