package dbcd

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAdminEndpoint 返回一个只有Admin角色才能访问的路由。
func (engine *Engine) GetAdminEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		resumeRequestBody(c)

		var response = NewRouterResponse()
		response.SetCodeAndMsg(0, "欢迎，Admin。")
		c.JSON(http.StatusOK, response)
	}
}
