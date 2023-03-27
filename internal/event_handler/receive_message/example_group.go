package receiveMessage

import (
	"xlab-feishu-robot/internal/pkg"

	_ "github.com/sirupsen/logrus"
)

func init(){
	groupMessageRegister(groupHelpMenu, "help")
}

func groupHelpMenu(messageevent *MessageEvent){
	pkg.Cli.MessageSend("chat_id",messageevent.Message.Chat_id,"text","this is a group test string")
}