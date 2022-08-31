package robotServer

import "github.com/gin-gonic/gin"

func Register(r *gin.Engine) {
	r.POST("/feishu_events", feishuEventHandler)
}
