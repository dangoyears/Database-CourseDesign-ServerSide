package dbcd

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetRegisterCourseEndpoint 返回“/register/course”处的路由。
func (engine *Engine) GetRegisterCourseEndpoint() gin.HandlerFunc {

	type registerCourseEndpointParam struct {
		CourseNumber  string `json:"courseId" form:"courseId" binding:"required"`
		StudentNumber string `json:"studentId" form:"studentId" binding:"required"`
	}

	return func(c *gin.Context) {
		var param registerCourseEndpointParam
		var response = NewRouterResponse()

		resumeRequestBody(c)
		if c.ShouldBind(&param) == nil {

			courseNumber, err := strconv.Atoi(param.CourseNumber)
			if err != nil {
				Trace(err)
			}

			studentNumber, err := strconv.Atoi(param.StudentNumber)
			if err != nil {
				Trace(err)
			}

			engine.AddCourseForStudent(courseNumber, studentNumber)
			response.SetCodeAndMsg(0, "成功注册学生课程信息。")
			c.JSON(http.StatusOK, response)
			return
		}
		response.SetCodeAndMsg(-1, "参数不足。必须提供非空的courseId和studentId。")
		c.JSON(http.StatusOK, response)
	}
}
