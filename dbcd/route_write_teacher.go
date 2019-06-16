package dbcd

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// GetWriteTeacherEndpoint 返回“/write/teacher”处的路由。
func (engine *Engine) GetWriteTeacherEndpoint() gin.HandlerFunc {

	type writeTeacherEndpointParam struct {
		CollegeName      string `form:"college" binding:"required"`
		Name             string `form:"name" binding:"required"`
		TeacherNumber    int    `form:"jobId" binding:"required"`
		Sex              string `form:"sex" binding:"required"`
		TeacherDegree    string `form:"education" binding:"required"`
		GraduationSchool string `form:"graduation" binding:"required"`
		Birthday         string `form:"birthday"`
		Identity         string `form:"idCard" binding:"required"`
		Password         string `form:"password" binding:"required"`
		Position         string `form:"status" binding:"required"`
	}

	return func(c *gin.Context) {
		var param writeTeacherEndpointParam
		var response = NewRouterResponse()

		if c.ShouldBind(&param) == nil {
			yearMonthAndDay := strings.Split(param.Birthday, "-")
			if len(yearMonthAndDay) != 3 {
				response.SetCodeAndMsg(1, "birthday格式不正确，必须为“yyyy-mm-dd”。")
				c.JSON(http.StatusOK, response)
				c.Abort()
				return
			}
			birthYear, _ := strconv.Atoi(yearMonthAndDay[0])
			birthMonth, _ := strconv.Atoi(yearMonthAndDay[1])
			birthDay, _ := strconv.Atoi(yearMonthAndDay[2])

			// TODO
			// 检查重复

			var human = Human{
				Name:         param.Name,
				Sex:          param.Sex,
				Birthday:     time.Date(birthYear, time.Month(birthMonth), birthDay, 0, 0, 0, 0, time.UTC),
				Identity:     param.Identity,
				PasswordHash: GeneratePasswordHash(param.Password),
			}

			engine.CreateTeacher(human, param.CollegeName, param.TeacherNumber, param.GraduationSchool, param.Position, param.TeacherDegree)
			response.SetCodeAndMsg(0, "此API尚不能检查参数错误以外的异常。如果一切顺利，教师将被创建。请使用/read/接口确认是否创建成功。信息修改未实现！由于提供了birthday参数，age参数已被忽略。")
			c.JSON(http.StatusOK, response)
			return
		}
		response.SetCodeAndMsg(-1, "参数不足。必须提供非空的college、name、jobId、sex、education、graduation、birthday、idCard、password和status参数。")
		c.JSON(http.StatusOK, response)
	}
}
