package receiveMessage

import (
	"xlab-feishu-robot/internal/pkg"

	_ "github.com/sirupsen/logrus"
)

func init(){
	p2pMessageRegister(p2pHelpMenu, "help")
}

func p2pHelpMenu(messageevent *MessageEvent){
	pkg.Cli.MessageSend("open_id",messageevent.Sender.Sender_id.Open_id,"text","this is a P2P test string")
}