package dbcd

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// GetWriteStudentEndpoint 返回“/write/student”处的路由。
func (engine *Engine) GetWriteStudentEndpoint() gin.HandlerFunc {

	type writeStudentEndpointParam struct {
		CollegeName            string `form:"college" binding:"required"`
		SpecialtyName          string `form:"specialty" binding:"required"`
		Grade                  int    `form:"grade" binding:"required"`
		ClassCode              int    `form:"class" binding:"required"`
		Name                   string `form:"name" binding:"required"`
		StudentNumber          int    `form:"studentId" binding:"required"`
		StudentDegreeAndStatus string `form:"status"`
		Sex                    string `form:"sex" binding:"required"`
		Birthday               string `form:"birthday" binding:"required"`
		Identity               string `form:"idCard" binding:"required"`
		Password               string `form:"password" binding:"required"`
		SchoolOfYear           int    `form:"yearSystem" binding:"required"`
	}

	return func(c *gin.Context) {
		var param writeStudentEndpointParam
		var response = NewRouterResponse()

		if BindContextIntoStruct(c, &param) == nil {
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

			engine.CreateStudent(human, param.CollegeName, param.SpecialtyName, param.Grade, param.ClassCode, param.StudentNumber)
			response.SetCodeAndMsg(0, "此API尚不能检查参数错误以外的异常。如果一切顺利，学生将被创建。status未实现！信息修改未实现！请使用/read/接口确认是否创建成功。由于提供了birthday参数，age参数已被忽略。")
			c.JSON(http.StatusOK, response)
			return
		}
		response.SetCodeAndMsg(-1, "参数不足。必须提供非空的college、specialty、grade、class、name、studentId、status、sex、birthday、idCard、password和yearSystem参数。")
		response["args"] = c.Params
		c.JSON(http.StatusOK, response)
	}
}
