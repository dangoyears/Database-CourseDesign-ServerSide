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
		CollegeName      string `json:"college" form:"college" binding:"required"`
		Name             string `json:"name" form:"name" binding:"required"`
		TeacherNumber    string `json:"jobId" form:"jobId" binding:"required"`
		Sex              string `json:"sex" form:"sex" binding:"required"`
		TeacherDegree    string `json:"education" form:"education" binding:"required"`
		GraduationSchool string `json:"graduation" form:"graduation" binding:"required"`
		Birthday         string `json:"birthday" form:"birthday"`
		Identity         string `json:"idCard" form:"idCard" binding:"required"`
		Password         string `json:"password" form:"password" binding:"required"`
		Position         string `json:"position" form:"position" binding:"required"`
	}

	return func(c *gin.Context) {
		var param writeTeacherEndpointParam
		var response = NewRouterResponse()

		resumeRequestBody(c)
		if c.ShouldBind(&param) == nil {
			yearMonthAndDay := strings.Split(param.Birthday, "-")
			if len(yearMonthAndDay) != 3 {
				response.SetCodeAndMsg(1, "birthday格式不正确，必须为“yyyy-mm-dd”。")
				c.JSON(http.StatusOK, response)
				c.Abort()
				return
			}
			birthYear, birthYearErr := strconv.Atoi(yearMonthAndDay[0])
			birthMonth, birthMonthErr := strconv.Atoi(yearMonthAndDay[1])
			birthDay, birthDayErr := strconv.Atoi(yearMonthAndDay[2])

			if birthYearErr != nil || birthMonthErr != nil || birthDayErr != nil {
				response.SetCodeAndMsg(1, "birthday格式不正确，必须为“yyyy-mm-dd”。")
				c.JSON(http.StatusOK, response)
				c.Abort()
				return
			}

			teacherNumber, teacherNumberErr := strconv.Atoi(param.TeacherNumber)
			if teacherNumberErr != nil {
				response.SetCodeAndMsg(1, "jobId必须能转换为整数。")
				c.JSON(http.StatusOK, response)
				c.Abort()
				return
			}

			position := param.Position
			if position != "教务办主任" && position != "普通教师" {
				response.SetCodeAndMsg(1, "status必须为教务办主任或普通教师。")
				c.JSON(http.StatusOK, response)
				c.Abort()
				return
			}

			var teacherInfo = TeacherInfo{
				CollegeName:      param.CollegeName,
				Name:             param.Name,
				Sex:              param.Sex,
				Birthday:         time.Date(birthYear, time.Month(birthMonth), birthDay, 0, 0, 0, 0, time.UTC),
				Identity:         param.Identity,
				PasswordHash:     GeneratePasswordHash(param.Password),
				TeacherNumber:    teacherNumber,
				GraduationSchool: param.GraduationSchool,
				Position:         position,
				TeacherDegree:    param.TeacherDegree,
			}

			if teacher := engine.GetTeacherByTeacherNumber(teacherNumber); teacher != nil {
				engine.UpdateTeacherAsInfo(teacherNumber, teacherInfo)
				return
			}

			engine.CreateTeacherAsInfo(teacherInfo)
			response.SetCodeAndMsg(0, "教师将被创建或更新。由于提供了birthday参数，age参数已被忽略。")
			c.JSON(http.StatusOK, response)
			return
		}
		response.SetCodeAndMsg(-1, "参数不足。必须提供非空的college、name、jobId、sex、education、graduation、birthday、idCard、password和status参数。")
		c.JSON(http.StatusOK, response)
	}
}
