package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary example custom controller
// @Tags custom_controller
// @Success 200 {string} OK
// @Router /api/example [post]
func Example(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
