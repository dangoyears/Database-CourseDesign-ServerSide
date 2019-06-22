package dbcd

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetWriteCollegeEndpoint 返回“/write/college”处的路由。
func (engine *Engine) GetWriteCollegeEndpoint() gin.HandlerFunc {

	type writeCollegeEndpointParam struct {
		CollegeName   string `json:"college" form:"college" binding:"required"`
		SpecialtyName string `json:"specialty" form:"specialty" binding:"required"`
		Grade         string `json:"grade" form:"grade" binding:"required"`
		ClassCode     string `json:"class" form:"class" binding:"required"`
	}

	return func(c *gin.Context) {
		param := GetArgs(c)
		var response = NewRouterResponse()

		collegeName, collegeNameOk := param["college"]
		specialtyName, specialtyNameOk := param["specialty"]
		grade, gradeOk := param["grade"]
		classCode, classCodeOk := param["class"]

		if collegeNameOk && specialtyNameOk && gradeOk && classCodeOk {

			engine.CreateClass(collegeName.(string), specialtyName.(string), grade.(int), classCode.(int))
			response.SetCodeAndMsg(0, "成功创建班级。")
			c.JSON(http.StatusOK, response)
			return

		}
		response.SetCodeAndMsg(-1, "参数不足。必须提供非空的college、specialty、grade和class参数。")
		c.JSON(http.StatusOK, response)
	}
}
