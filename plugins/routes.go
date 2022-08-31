package plugins

import (
	exampleController "xlab-feishu-robot/plugins/controller/example"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	r.POST("/api/example", exampleController.Controller)
}
