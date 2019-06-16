package dbcd

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetDeleteClassEndpoint 提供“/delte/class”的路由。
func (engine *Engine) GetDeleteClassEndpoint() gin.HandlerFunc {

	type deleteClassEndpointParam struct {
		CollegeName   string `form:"college"`
		SpecialtyName string `form:"specialty" binding:"required"`
		Grade         int    `form:"grade" binding:"required"`
		ClassCode     int    `form:"class" binding:"required"`
	}

	return func(c *gin.Context) {
		var response = NewRouterResponse()
		var param deleteClassEndpointParam

		if c.ShouldBind(&param) == nil {
			engine.DeleteClassBySpecialtyNameGradeAndCode(param.SpecialtyName, param.Grade, param.ClassCode)

			response.SetCodeAndMsg(0, "如无意外，指定的班级将被删除。")
			c.JSON(http.StatusOK, response)
			return
		}

		// 参数不足
		response.SetCodeAndMsg(-1, "参数不足。必须提供非空的specialty、grade和class参数。")
		c.JSON(http.StatusOK, response)
	}
}
