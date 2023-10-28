package receiveMessage

import (
	_ "github.com/sirupsen/logrus"
)

func init() {
	groupMessageRegister(groupHelpMenu, "help")
}

func groupHelpMenu(messageevent *MessageEvent) {
	SendMessage(GroupChatId, messageevent.Message.Chat_id, Text, "this is a group test string")
}
