package dbcd

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type logoffParam struct {
	Token string `form:"token"`
}

func (engine *Engine) getLogoutEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		var response = NewRouterResponse()
		var param logoffParam

		if c.ShouldBind(&param) == nil {
			token := param.Token

			engine.keeper.Logoff(token)

			response.setCodeAndMsg(0, "成功退出登陆。")
			c.JSON(http.StatusOK, response)
			return
		}

		// 参数不足
		response.setCodeAndMsg(-1, "参数不足。必须提供token参数。")
		c.JSON(http.StatusOK, response)
	}
}
