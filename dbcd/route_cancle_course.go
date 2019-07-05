package dbcd

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetCancelCourseEndpoint 提供“/cancel/course”的路由。
func (engine *Engine) GetCancelCourseEndpoint() gin.HandlerFunc {

	type cancelCourseEndpointParam struct {
		StudentNumber string `json:"studentId" form:"studentId" binding:"required"`
		CourseNumber  string `json:"courseId" form:"courseId" binding:"required"`
	}

	return func(c *gin.Context) {
		resumeRequestBody(c)

		var response = NewRouterResponse()
		var param cancelCourseEndpointParam

		if c.ShouldBind(&param) == nil {
			studentNumber, studentNumberErr := strconv.Atoi(param.StudentNumber)
			courseNumber, courseNumberErr := strconv.Atoi(param.CourseNumber)

			if courseNumberErr != nil || studentNumberErr != nil {
				response.SetCodeAndMsg(-1, "id必须能转换为整数。")
				c.JSON(http.StatusOK, response)
				return
			}

			engine.RemoveCourseForStudent(courseNumber, studentNumber)
			response.SetCodeAndMsg(0, "指定student的course已删除。")
			c.JSON(http.StatusOK, response)
			return
		}

		// 参数不足
		response.SetCodeAndMsg(-1, "参数不足。必须提供非空的id参数。")
		c.JSON(http.StatusOK, response)
	}
}
