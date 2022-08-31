package plugins

import (
	customRobots "xlab-feishu-robot/plugins/custom_robots"
	feishuEventHandler "xlab-feishu-robot/plugins/feishu_event_handler"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	customRobots.Init(r)
	feishuEventHandler.Init()
}
