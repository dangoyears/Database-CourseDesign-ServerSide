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
		CollegeName            string `json:"college" form:"college" binding:"required"`
		SpecialtyName          string `json:"specialty" form:"specialty" binding:"required"`
		Grade                  string `json:"grade" form:"grade" binding:"required"`
		ClassCode              string `json:"class" form:"class" binding:"required"`
		Name                   string `json:"name" form:"name" binding:"required"`
		StudentNumber          string `json:"studentId" form:"studentId" binding:"required"`
		StudentDegreeAndStatus string `json:"status" form:"status"`
		Sex                    string `json:"sex" form:"sex" binding:"required"`
		Birthday               string `json:"birthday" form:"birthday" binding:"required"`
		Identity               string `json:"idCard" form:"idCard" binding:"required"`
		Password               string `json:"password" form:"password" binding:"required"`
		YearOfSchool           string `json:"yearSystem" form:"yearSystem" binding:"required"`
	}

	return func(c *gin.Context) {
		var param writeStudentEndpointParam
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

			collegeName := param.CollegeName
			specialtyName := param.SpecialtyName
			grade, gradeOK := strconv.Atoi(param.Grade)
			classCode, classCodeOK := strconv.Atoi(param.ClassCode)

			if gradeOK != nil || classCodeOK != nil {
				response.SetCodeAndMsg(0, "grade参数或class参数不是有效的形式。")
				c.JSON(http.StatusOK, response)
				return
			}

			studentNumber, studentNumberErr := strconv.Atoi(param.StudentNumber)
			if studentNumberErr != nil {
				response.SetCodeAndMsg(0, "studentId参数必须可转换成整数。")
				c.JSON(http.StatusOK, response)
				return
			}

			yearOfSchool, yearOfSchoolErr := strconv.Atoi(param.YearOfSchool)
			if yearOfSchoolErr != nil {
				response.SetCodeAndMsg(0, "yearSystem参数必须可转换成整数。")
				c.JSON(http.StatusOK, response)
				return
			}

			var (
				studentDegree string
				status        string
			)

			switch param.StudentDegreeAndStatus {
			case "在读本科生":
				studentDegree, status = "学士", "在读"
			case "毕业本科生":
				studentDegree, status = "学士", "毕业"
			case "在读研究生":
				studentDegree, status = "硕士", "在读"
			case "毕业研究生":
				studentDegree, status = "硕士", "毕业"
			default:
				response.SetCodeAndMsg(0, "status参数无效，参数值必须为在读本科生、毕业本科生、在读研究生、毕业研究生其一。")
				c.JSON(http.StatusOK, response)
				return
			}

			var studentInfo = StudentInfo{
				CollegeName:   collegeName,
				SpecialtyName: specialtyName,
				Grade:         grade,
				ClassCode:     classCode,
				Name:          param.Name,
				Sex:           param.Sex,
				Birthday:      time.Date(birthYear, time.Month(birthMonth), birthDay, 0, 0, 0, 0, time.UTC),
				Identity:      param.Identity,
				PasswordHash:  GeneratePasswordHash(param.Password),
				StudentNumber: studentNumber,
				YearOfSchool:  yearOfSchool,
				StudentDegree: studentDegree,
				Status:        status,
			}

			// 如果Identity与现有学生重复，则进行更新，否则创建
			if student := engine.GetStudentByStudentNubmer(studentNumber); student != nil {
				engine.UpdateStudentAsInfo(studentNumber, studentInfo)
				response.SetCodeAndMsg(0, "学生信息将更新。由于提供了birthday参数，age参数已被忽略。")
				c.JSON(http.StatusOK, response)
				return
			}

			engine.CreateStudentAsInfo(studentInfo)
			response.SetCodeAndMsg(0, "学生将被创建。由于提供了birthday参数，age参数已被忽略。")
			c.JSON(http.StatusOK, response)
			return
		}
		response.SetCodeAndMsg(-1, "参数不足。必须提供非空的college、specialty、grade、class、name、studentId、status、sex、birthday、idCard、password和yearSystem参数。")
		response["args"] = c.Params
		c.JSON(http.StatusOK, response)
	}
}
