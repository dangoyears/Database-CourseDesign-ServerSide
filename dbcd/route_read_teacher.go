package dbcd

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetReadTeacherEndpoint 返回“/read/teacher”处的路由。
// @未完成
func (engine *Engine) GetReadTeacherEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		var response = NewRouterResponse()

		teacherInfo := engine.GetTeacherInfo()
		response["data"] = teacherInfo
		response.SetCodeAndMsg(0, "已读取。")
		c.JSON(http.StatusAccepted, response)
	}
}
