package receiveMessage

import (
	"context"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
	"github.com/sirupsen/logrus"
)

// Receive dispatch message according to chat type
func Receive(_ context.Context, event *larkim.P2MessageReceiveV1) error {
	chatType := *event.Event.Message.ChatType
	switch chatType {
	case "p2p":
		p2p(event)
	case "group":
		group(event)
	// add more chat type here if needed

	default:
		logrus.WithFields(logrus.Fields{"chat type": chatType}).Warn("Receive message, but this chat type is not supported")
	}

	return nil
}
