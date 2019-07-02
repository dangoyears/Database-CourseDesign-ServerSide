package dbcd

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetDeleteCourseEndpoint 提供“/delte/course”的路由。
func (engine *Engine) GetDeleteCourseEndpoint() gin.HandlerFunc {

	type deleteCourseEndpointParam struct {
		CourseNumber string `json:"id" form:"id" binding:"required"`
	}

	return func(c *gin.Context) {
		resumeRequestBody(c)

		var response = NewRouterResponse()
		var param deleteCourseEndpointParam

		if c.ShouldBind(&param) == nil {
			courseNumber, courseNumberErr := strconv.Atoi(param.CourseNumber)

			if courseNumberErr != nil {
				response.SetCodeAndMsg(-1, "id必须能转换为整数。")
				c.JSON(http.StatusOK, response)
				return
			}

			engine.DeleteCourseByCourseNumber(courseNumber)
			response.SetCodeAndMsg(0, "指定course已删除。")
			c.JSON(http.StatusOK, response)
			return
		}

		// 参数不足
		response.SetCodeAndMsg(-1, "参数不足。必须提供非空的id参数。")
		c.JSON(http.StatusOK, response)
	}
}
