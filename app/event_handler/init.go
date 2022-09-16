package eventHandler

import (
	"xlab-feishu-robot/pkg/dispatcher"
	"xlab-feishu-robot/app/event_handler/receive_message"
)


func Init() {
	// register your handlers here
	// example
	dispatcher.RegisterListener(receiveMessage.Receive, "im.message.receive_v1")
}
