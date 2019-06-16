package dbcd

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetReadCollegeEndpoint 返回“/read/college”处的路由。
func (engine *Engine) GetReadCollegeEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		var response = NewRouterResponse()

		classInfo := engine.GetClassInfo()
		response["data"] = classInfo
		response.SetCodeAndMsg(0, "已读取。")
		c.JSON(http.StatusAccepted, response)
	}
}
