package dbcd

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type loginStatusParam struct {
	Token string `form:"token"`
}

func (engine *Engine) getLoginStatusEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		var response = NewRouterResponse()
		var param loginStatusParam

		if c.ShouldBind(&param) == nil {
			token := param.Token
			engine.keeper.Logoff(token)

			loginType, _ := engine.keeper.getLoginType(token)

			response.Data["role"]
			response.setCodeAndMsg(0, "已查询。")
			c.JSON(http.StatusOK, response)
			return
		}

		response.Data["role"] = "anonymous"
		response.setCodeAndMsg(0, "未登录")
		c.JSON(http.StatusOK, response)
	}
}
