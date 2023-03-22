package app

import (
	eventHandler "xlab-feishu-robot/app/event_handler"
	"xlab-feishu-robot/app/router"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	eventHandler.Init()
	router.Register(r)
}
