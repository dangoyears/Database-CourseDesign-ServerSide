package dbcd

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (engine *Engine) getEchoRoute() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, GetArgs(c))
	}
}
