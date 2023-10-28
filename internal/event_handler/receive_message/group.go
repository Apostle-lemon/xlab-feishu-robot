package receiveMessage

import (
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
	"github.com/sirupsen/logrus"
	"strings"
)

func group(event *larkim.P2MessageReceiveV1) {
	messageType := *event.Event.Message.MessageType
	switch strings.ToUpper(messageType) {
	case "TEXT":
		groupTextMessage(event)
	default:
		logrus.WithFields(logrus.Fields{"message type": messageType}).Warn("Receive group message, but this type is not supported")
	}
}

func groupTextMessage(event *larkim.P2MessageReceiveV1) {
	// get the pure text message
	content := *event.Event.Message.Content
	content = strings.TrimSuffix(strings.TrimPrefix(content, "{\"text\":\""), "\"}")
	// 在群组中，消息内容的前面往往会有一个@机器人的字符串，需要去掉
	content = content[strings.Index(content, " ")+1:]
	event.Event.Message.Content = &content
	logrus.WithFields(logrus.Fields{"message content": content}).Info("Receive group TEXT message")

	switch content {
	case "help":
		groupHelpMenu(event)
	default:
		logrus.WithFields(logrus.Fields{"message content": content}).Warn("Receive group TEXT message, but this content does not have a handler")
	}
}
