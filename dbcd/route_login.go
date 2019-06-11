package dbcd

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type loginEndpointParam struct {
	Type string `form:"type" binding:"required"` // {"anonymous", "admin", "student", "teacher"}之一
	Name string `form:"name" binding:"required"`
	Pass string `form:"pass" binding:"required"`
}

// GetLoginEndpoint 提供“/login”的路由。
func (engine *Engine) GetLoginEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		var response = NewRouterResponse()
		var param loginEndpointParam

		if c.ShouldBind(&param) == nil {
			name, pass := param.Name, param.Pass
			var token string

			switch param.Type {
			case "admin":
				token = engine.LoginAdmin(name, pass)
			case "student":
				token = engine.LoginStudent(name, pass)
			case "teacher":
				token = engine.LoginTeacher(name, pass)
			}
			response["token"] = token

			if token != "" {
				response.SetCodeAndMsg(0, "认证成功。")
				c.JSON(http.StatusOK, response)
				return
			}

			response.SetCodeAndMsg(1, "认证失败。type、name或pass错误，或者用户不存在。")
			c.JSON(http.StatusOK, response)
			return
		}

		// 参数不足
		response.SetCodeAndMsg(-1, "参数不足。必须提供非空的type、name和pass参数。")
		c.JSON(http.StatusOK, response)
	}
}
