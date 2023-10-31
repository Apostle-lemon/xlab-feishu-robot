package internal

import (
	"xlab-feishu-robot/internal/router"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {
	router.Register(r)
}
