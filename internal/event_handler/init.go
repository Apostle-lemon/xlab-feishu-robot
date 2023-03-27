package eventHandler

import (
	"xlab-feishu-robot/internal/dispatcher"
	receiveMessage "xlab-feishu-robot/internal/event_handler/receive_message"
)

// bind event handlers with event types to dispatcher
func Init() {
	dispatcher.RegisterListener(receiveMessage.Receive, "im.message.receive_v1")
}