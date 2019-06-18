package dbcd

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetWelcomeEndpoint 返回后端接口欢迎路由
func (engine *Engine) GetWelcomeEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		var response = NewRouterResponse()
		response.SetCodeAndMsg(0, `欢迎使用DBCD后端接口。\(@^0^@)/`)
		response["doc"] = "https://github.com/dangoyears/Database-CourseDesign-Docs"
		response["ver"] = "0.0.3"
		c.JSON(http.StatusOK, response)
	}
}
