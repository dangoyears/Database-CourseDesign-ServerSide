package dbcd

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetRoleEndpoint 提供“/role”的路由。
func (engine *Engine) GetRoleEndpoint() gin.HandlerFunc {

	type roleEndpointParam struct {
		Token string `form:"token" binding:"required"`
	}

	return func(c *gin.Context) {
		var response = NewRouterResponse()
		var param roleEndpointParam

		if BindContextIntoStruct(c, &param) == nil {
			token := param.Token

			role := engine.keeper.GetRole(token)

			response["role"] = role
			response.SetCodeAndMsg(0, "已查询。")
			c.JSON(http.StatusOK, response)
			return
		}

		response["role"] = "anonymous"
		response.SetCodeAndMsg(-1, "未传入token参数，或传入token为空。")
		c.JSON(http.StatusOK, response)
	}
}
