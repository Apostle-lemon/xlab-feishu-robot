package router

import (
	"xlab-feishu-robot/app/dispatcher"

	"github.com/gin-gonic/gin"
)

func RegisterDispatcher(r *gin.Engine) {
	r.POST("/feishu_events", dispatcher.Dispatcher)
}

func RegisterPOST(r *gin.Engine, relativePath string, handlers ...gin.HandlerFunc) {
	r.POST(relativePath, handlers...)
}

func RegisterGET(r *gin.Engine, relativePath string, handlers ...gin.HandlerFunc) {
	r.GET(relativePath, handlers...)
}

// add PATCH, DELETE etc. if needed
