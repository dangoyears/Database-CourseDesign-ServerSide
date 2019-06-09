package dbcd

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type loginParam struct {
	username string `form:"name"`
	password string `form:"pass"`
	userType string `form:"type"`
}

// LoginEndpoint 提供“/login”的路由
func LoginEndpoint(c *gin.Context) {
	var response gin.H
	var param loginParam

	if c.ShouldBind(&param) == nil {
		if param.userType == "admin" && param.username == "admin" && param.password == "dangoyears" {
			// 管理员成功登陆
			c.JSON(http.StatusOK, response)
		}
	}

	// 参数不足
	response["msg"] = "必须提供name和pass参数。"
	c.JSON(http.StatusOK, response)
}
