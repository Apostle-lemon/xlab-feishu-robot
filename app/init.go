package app

import (
	"xlab-feishu-robot/app/router"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	router.RegisterDispatcher(r)
}
