package example

import "github.com/gin-gonic/gin"

func Controller(c *gin.Context) {
	c.String(200, "OK")
}
