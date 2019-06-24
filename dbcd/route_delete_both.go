package dbcd

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetDeleteBothEndpoint 提供“/delte/both”的路由。
func (engine *Engine) GetDeleteBothEndpoint() gin.HandlerFunc {

	type deleteBothEndpointParam struct {
		Role                   string `json:"type" form:"type" binding:"required"`
		TeacherOrStudentNumber string `json:"id" form:"id" binding:"required"`
	}

	return func(c *gin.Context) {
		resumeRequestBody(c)

		var response = NewRouterResponse()
		var param deleteBothEndpointParam

		if c.ShouldBind(&param) == nil {
			teacherOrStudentNumber, teacherOrStudentNumberErr := strconv.Atoi(param.TeacherOrStudentNumber)
			if teacherOrStudentNumberErr != nil {
				response.SetCodeAndMsg(1, "id必须可转换为数字。")

				c.JSON(http.StatusOK, response)
				return
			}

			switch param.Role {
			case "teacher":
				engine.DeleteTeacherByTeacherNumber(teacherOrStudentNumber)
				response.SetCodeAndMsg(0, "如无意外，指定的教师将被删除。")
			case "student":
				engine.DeleteStudentByStudentNubmer(teacherOrStudentNumber)
				response.SetCodeAndMsg(0, "如无意外，指定的学生将被删除。")
			default:
				response.SetCodeAndMsg(1, "不合理的type："+param.Role)
			}

			c.JSON(http.StatusOK, response)
			return
		}

		// 参数不足
		response.SetCodeAndMsg(-1, "参数不足。必须提供非空的type和id参数。")
		c.JSON(http.StatusOK, response)
	}
}
