package receiveMessage

import (
	_ "github.com/sirupsen/logrus"
)

func init() {
	p2pMessageRegister(p2pHelpMenu, "help")
}

func p2pHelpMenu(messageevent *MessageEvent) {
	SendMessage(UserOpenId, messageevent.Sender.Sender_id.Open_id, Text, "this is a p2p test string")
}
