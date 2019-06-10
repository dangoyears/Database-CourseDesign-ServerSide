package dbcd

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type loginParam struct {
	Type string `form:"type"` // {"admin", "student", "teacher"}之一
	Name string `form:"name"`
	Pass string `form:"pass"`
}

// GetLoginEndpoint 提供“/login”的路由
func (engine *Engine) getLoginEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		var response = NewRouterResponse()
		var param loginParam

		if c.ShouldBind(&param) == nil {
			name, pass := param.Name, param.Pass
			var token string

			switch param.Type {
			case "admin":
				token = engine.keeper.LoginAdmin(name, pass)
			case "student":
				token = engine.keeper.LoginStudent(name, pass)
			case "teacher":
				token = engine.keeper.LoginTeacher(name, pass)
			}
			response.Data["token"] = token

			if token != "" {
				response.setCodeAndMsg(0, "认证成功。")
				c.JSON(http.StatusOK, response)
				return
			}

			response.setCodeAndMsg(1, "认证失败。type、name或pass错误，或者用户不存在。")
			c.JSON(http.StatusOK, response)
			return
		}

		// 参数不足
		response.setCodeAndMsg(-1, "参数不足。必须提供type、name和pass参数。")
		c.JSON(http.StatusOK, response)
	}
}
