package dbcd

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetLogoutEndpoint 提供“/logout”的路由。
func (engine *Engine) GetLogoutEndpoint() gin.HandlerFunc {

	type logoutEndpointParam struct {
		Token string `form:"token" binding:"required"`
	}

	return func(c *gin.Context) {
		var response = NewRouterResponse()
		var param logoutEndpointParam

		if c.ShouldBind(&param) == nil {
			token := param.Token

			engine.keeper.RemoveToken(token)

			response.SetCodeAndMsg(0, "token已失效。")
			c.JSON(http.StatusOK, response)
			return
		}

		// 参数不足
		response.SetCodeAndMsg(-1, "参数不足。必须提供非空的token参数。")
		c.JSON(http.StatusOK, response)
	}
}
