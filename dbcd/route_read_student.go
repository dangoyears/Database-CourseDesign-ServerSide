package dbcd

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetReadStudentEndpoint 返回“/read/student”处的路由。
// @未完成
func (engine *Engine) GetReadStudentEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		var response = NewRouterResponse()

		studentInfo := engine.GetStudentInfo()
		response["data"] = studentInfo
		response.SetCodeAndMsg(0, "已读取。")
		c.JSON(http.StatusAccepted, response)
	}
}
