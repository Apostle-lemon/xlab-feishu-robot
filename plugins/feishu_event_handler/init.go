package feishuEventHandler

import (
	"xlab-feishu-robot/app/dispatcher"
	"xlab-feishu-robot/plugins/feishu_event_handler/example"
)

func Init() {
	// register your handlers here

	// example
	dispatcher.RegisterListener(example.Handler, "example")
}
