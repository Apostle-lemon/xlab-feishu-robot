package receiveMessage

import (
	"github.com/sirupsen/logrus"
	"strings"
)

var p2pMessageMap = make(map[string]messageHandler)

func p2p(messageevent *MessageEvent){
	switch strings.ToUpper(messageevent.Message.Message_type) {
	case "TEXT":
		p2pTextMessage(messageevent)
	default:
		logrus.WithFields(logrus.Fields{"message type": messageevent.Message.Message_type}).Warn("Receive p2p message, but this type is not supported")
	}
}

func p2pTextMessage(messageevent *MessageEvent){
	// get the pure text message
	messageevent.Message.Content = strings.TrimSuffix(strings.TrimPrefix(messageevent.Message.Content, "{\"text\":\""), "\"}")
	logrus.WithFields(logrus.Fields{"message content": messageevent.Message.Content}).Info("Receive p2p TEXT message")

	if handler, exists := p2pMessageMap[messageevent.Message.Content]; exists {
		handler(messageevent)
		return
	} else {
		logrus.Error("p2p message failed to find event handler: ", messageevent.Message.Content)
		return
	}
}

func p2pMessageRegister(f messageHandler, s string) {
	
	if _, isEventExist := p2pMessageMap[s]; isEventExist {
		logrus.Warning("Double declaration of group message handler: ", s)
	}
	p2pMessageMap[s] = f

}