package internal

import (
	eventHandler "xlab-feishu-robot/internal/event_handler"
	"xlab-feishu-robot/internal/router"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	eventHandler.Init()
	router.Register(r)
}
