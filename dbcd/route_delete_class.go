package dbcd

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetDeleteClassEndpoint 提供“/delte/class”的路由。
func (engine *Engine) GetDeleteClassEndpoint() gin.HandlerFunc {

	type deleteClassEndpointParam struct {
		CollegeName   string `json:"college" form:"college"`
		SpecialtyName string `json:"specialty" form:"specialty" binding:"required"`
		Grade         string `json:"grade" form:"grade" binding:"required"`
		ClassCode     string `json:"class" form:"class" binding:"required"`
		Sum           string `json:"sum" form:"sum"`
	}

	return func(c *gin.Context) {
		resumeRequestBody(c)

		var response = NewRouterResponse()
		var param deleteClassEndpointParam

		if c.ShouldBind(&param) == nil {
			specialtyName := param.SpecialtyName
			grade, gradeOK := strconv.Atoi(param.Grade)
			classCode, classCodeOK := strconv.Atoi(param.ClassCode)

			if gradeOK != nil || classCodeOK != nil {
				response.SetCodeAndMsg(0, "grade参数或class参数不是有效的形式。")
				c.JSON(http.StatusOK, response)
				return
			}

			engine.DeleteClassBySpecialtyNameGradeAndCode(specialtyName, grade, classCode)

			response.SetCodeAndMsg(0, "如无意外，指定的班级"+specialtyName+strconv.Itoa(grade)+strconv.Itoa(classCode)+"将被删除。sum参数已被忽略。")
			c.JSON(http.StatusOK, response)
			return
		}

		// 参数不足
		response.SetCodeAndMsg(-1, "参数不足。必须提供非空的specialty、grade和class参数。")
		c.JSON(http.StatusOK, response)
	}
}
