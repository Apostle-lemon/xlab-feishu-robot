package receiveMessage

import (
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

func groupHelpMenu(event *larkim.P2MessageReceiveV1) {
	SendMessage(GroupChatId, *event.Event.Message.ChatId, Text, "this is a group test string")
}
