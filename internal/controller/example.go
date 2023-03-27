package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Example(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
