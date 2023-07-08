package healthcheck

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthcheckController struct {
}

func NewHealthcheckController() *HealthcheckController {
	return &HealthcheckController{}
}

func (con *HealthcheckController) Livez(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}
