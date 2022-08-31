package customRobots

import (
	"xlab-feishu-robot/app/router"
	"xlab-feishu-robot/plugins/custom_robots/example"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	// register your plugin's router here

	// example
	router.RegisterPlugin(r, "/api/example", example.Controller)
}
