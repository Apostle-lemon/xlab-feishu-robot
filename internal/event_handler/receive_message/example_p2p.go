package receiveMessage

import (
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

func p2pHelpMenu(event *larkim.P2MessageReceiveV1) {
	SendMessage(UserOpenId, *event.Event.Sender.SenderId.OpenId, Text, "this is a p2p test string")
}
