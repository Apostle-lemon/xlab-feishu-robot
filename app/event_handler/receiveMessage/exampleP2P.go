package receiveMessage

import (
	"xlab-feishu-robot/pkg/global"
	_ "github.com/sirupsen/logrus"
)

func init(){
	p2pMessageRegister(p2pHelpMenu, "help")
}

func p2pHelpMenu(messageevent *MessageEvent){
	global.Cli.Send("open_id",messageevent.Sender.Sender_id.Open_id,"text","this is a P2P test string")
}