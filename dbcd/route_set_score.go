package dbcd

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetSetScoreEndpoint 返回“/set/score”处的路由。
func (engine *Engine) GetSetScoreEndpoint() gin.HandlerFunc {

	type setScoreEndpointParam struct {
		CourseNumber            string            `json:"courseId" binding:"required"`
		StudentNumbersAndScores map[string]string `json:"students" binding:"required"`
	}

	return func(c *gin.Context) {
		var param setScoreEndpointParam
		var response = NewRouterResponse()

		resumeRequestBody(c)
		if c.ShouldBind(&param) == nil {
			courseNumber, courseNumberErr := strconv.Atoi(param.CourseNumber)

			if courseNumberErr != nil {
				response.SetCodeAndMsg(-1, "courseId必须能转换为整数。")
				c.JSON(http.StatusOK, response)
				return
			}

			for studentNumberStr, scoreStr := range param.StudentNumbersAndScores {
				studentNumber, studentNumberErr := strconv.Atoi(studentNumberStr)
				score, scoreErr := strconv.Atoi(scoreStr)

				if studentNumberErr != nil || scoreErr != nil {
					response.SetCodeAndMsg(-1, "学号及成绩必须能转换为整数。")
					c.JSON(http.StatusOK, response)
					return
				}
				engine.UpdateCourseStudentScore(courseNumber, studentNumber, score)
			}

			response.SetCodeAndMsg(0, "成功更新成绩。")
			c.JSON(http.StatusOK, response)
			return
		}
		response.SetCodeAndMsg(-1, "参数不足。必须提供非空的courseId和students参数。")
		c.JSON(http.StatusOK, response)
	}
}
