package dbcd

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetDeleteBothEndpoint 提供“/delte/both”的路由。
func (engine *Engine) GetDeleteBothEndpoint() gin.HandlerFunc {

	type deleteBothEndpointParam struct {
		Role                   string `form:"type" binding:"required"`
		TeacherOrStudentNumber int    `form:"id" binding:"required"`
	}

	return func(c *gin.Context) {
		var response = NewRouterResponse()
		var param deleteBothEndpointParam

		if BindContextIntoStruct(c, &param) == nil {
			switch param.Role {
			case "teacher":
				engine.DeleteTeacherByTeacherNumber(param.TeacherOrStudentNumber)
			case "student":
				engine.DeleteStudentByStudentNubmer(param.TeacherOrStudentNumber)
			default:
				response.SetCodeAndMsg(1, "不合理的type："+param.Role)
				c.JSON(http.StatusOK, response)
				c.Abort()
				return
			}

			response.SetCodeAndMsg(0, "如无意外，指定的教师/学生将被删除。")
			c.JSON(http.StatusOK, response)
			return
		}

		// 参数不足
		response.SetCodeAndMsg(-1, "参数不足。必须提供非空的type和id参数。")
		c.JSON(http.StatusOK, response)
	}
}
