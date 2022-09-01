package eventHandler

import "xlab-feishu-robot/pkg/dispatcher"

func Init() {
	// register your handlers here
	// example
	dispatcher.RegisterListener(example, "im.message.receive_v1")
}
