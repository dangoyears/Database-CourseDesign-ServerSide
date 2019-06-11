package dbcd

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type loginStatusParam struct {
	Token string `form:"token" binding:"required"`
}

// GetStatusLoginEndpoint 提供“/status/login”的路由。
func (engine *Engine) GetStatusLoginEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		var response = NewRouterResponse()
		var param loginStatusParam

		if c.ShouldBind(&param) == nil {
			token := param.Token

			role := engine.keeper.GetRole(token)

			response.Data["role"] = role
			response.SetCodeAndMsg(0, "已查询。")
			c.JSON(http.StatusOK, response)
			return
		}

		response.Data["role"] = "anonymous"
		response.SetCodeAndMsg(-1, "未传入token参数，或传入token为空。")
		c.JSON(http.StatusOK, response)
	}
}
