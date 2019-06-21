package dbcd

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetLoginEndpoint 提供“/login”的路由。
func (engine *Engine) GetLoginEndpoint() gin.HandlerFunc {

	type loginEndpointParam struct {
		Type     string `form:"type" binding:"required"`
		Username string `form:"user" binding:"required"`
		Password string `form:"pass" binding:"required"`
	}

	return func(c *gin.Context) {
		var response = NewRouterResponse()
		param := GetArgs(c)

		name, nameOk := param["user"].(string)
		pass, passOK := param["pass"].(string)
		usertype, typeOK := param["type"].(string)

		log.Println(name, pass, usertype)

		if nameOk && passOK && typeOK && name != "" && pass != "" {
			var token string

			switch usertype { // {"anonymous", "admin", "student", "teacher"}之一
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

			response.SetCodeAndMsg(1, "认证失败。type、user或pass错误，或者用户不存在。")
			c.JSON(http.StatusOK, response)
			return
		}

		// 参数不足
		response.SetCodeAndMsg(-1, "参数不足。必须提供非空的type、user和pass参数。")
		c.JSON(http.StatusOK, response)
	}
}
