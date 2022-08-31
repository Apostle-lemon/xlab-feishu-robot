package robotServer

import exampleHandler "xlab-feishu-robot/robot_server/handlers/example"

func init() {
	registerListener(exampleHandler.Handler, "example_handler")

	// register your listeners here
}
