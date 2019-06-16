package dbcd

import (
	"net/http"
	"strconv"

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
		var param writeCollegeEndpointParam
		var response = NewRouterResponse()

		if err := c.ShouldBind(&param); err == nil {
			collegeName, specailtyName := param.CollegeName, param.SpecialtyName

			grade, err := strconv.Atoi(param.Grade)
			if err != nil {
				engine.Trace(err)
			}

			classCode, err := strconv.Atoi(param.ClassCode)
			if err != nil {
				engine.Trace(err)
			}

			engine.CreateClass(collegeName, specailtyName, grade, classCode)
			response.SetCodeAndMsg(0, "成功创建班级。")
			c.JSON(http.StatusOK, response)
			return
		} else {
			engine.Trace(err)
		}
		response.SetCodeAndMsg(-1, "参数不足。必须提供非空的college、specialty、grade和class参数。")
		c.JSON(http.StatusOK, response)
	}
}
