package router

import (
	"xlab-feishu-robot/app/controller"
	"xlab-feishu-robot/pkg/dispatcher"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	// register your controllers here
	// example
	r.POST("/api/example", controller.Example)

	// DO NOT CHANGE LINES BELOW
	// register dispatcher
	r.POST("/feiShu/Event", dispatcher.Dispatcher)
}
