package example

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary example custom controller
// @Tags custom_robots
// @Success 200 {string} OK
// @Router /api/example [post]
func Controller(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
