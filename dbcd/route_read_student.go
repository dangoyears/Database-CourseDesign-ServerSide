package dbcd

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// GetReadStudentEndpoint 返回“/read/student”处的路由。
func (engine *Engine) GetReadStudentEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		var response = NewRouterResponse()

		response["data"] = []gin.H{}

		students := engine.GetStudentInfo()

		for _, student := range students {
			studentInfo := make(gin.H)

			studentInfo["college"] = student.CollegeName
			studentInfo["specialty"] = student.SpecialtyName
			studentInfo["grade"] = strconv.Itoa(student.Grade)
			studentInfo["class"] = strconv.Itoa(student.ClassCode)
			studentInfo["name"] = student.Name
			studentInfo["sex"] = student.Sex
			studentInfo["yearSystem"] = strconv.Itoa(student.YearOfSchool)
			studentInfo["studentId"] = fmt.Sprintf("%010d", student.StudentNumber)

			studentInfo["status"] = student.Status
			switch student.StudentDegree {
			case "学士":
				studentInfo["status"] = studentInfo["status"].(string) + "本科生"
			case "硕士":
				studentInfo["status"] = studentInfo["status"].(string) + "研究生"
			case "博士":
				studentInfo["status"] = studentInfo["status"].(string) + "博士生"
			}

			studentInfo["birthday"] = fmt.Sprintf("%d-%d-%d", student.Birthday.Year(), student.Birthday.Month(), student.Birthday.Day())
			studentInfo["age"] = fmt.Sprintf("%d", time.Now().Year()-student.Birthday.Year())
			studentInfo["idCard"] = student.Identity
			studentInfo["schedule"] = []gin.H{}

			courses := engine.GetCourseByStudentNumber(student.StudentNumber)
			for _, course := range courses {
				courseInfo := make(gin.H)

				courseInfo["name"] = course.CourseName
				courseInfo["id"] = fmt.Sprintf("%010d", course.CourseNumber)
				courseInfo["credit"] = strconv.Itoa(course.Credits)
				switch course.CourseProperty {
				case 1:
					courseInfo["nature"] = "专业必修课"
				case 2:
					courseInfo["nature"] = "专业选修课"
				case 3:
					courseInfo["nature"] = "通识性选修课"
				case 4:
					courseInfo["nature"] = "体育选修课"
				}
				courseInfo["accommodate"] = strconv.Itoa(course.Accommodate)
				courseInfo["selectedSum"] = strconv.Itoa(engine.GetCourseSelectionSumByCourseNumber(course.CourseNumber))
				courseInfo["time"] = course.Time

				courseInfo["teachers"] = []string{}
				teacherHumanIDs := engine.GetTeacherHumanIDsByCourseID(course.CourseID)
				for _, humanID := range teacherHumanIDs {
					courseInfo["teachers"] = append(courseInfo["teachers"].([]string), engine.GetHumanByID(humanID).Name)
				}

				courseInfo["courseLeader"] = ""
				if course.LeadTeacherHumanID != nil {
					courseInfo["courseLeader"] = engine.GetHumanByID(*course.LeadTeacherHumanID).Name
				}

				courseInfo["address"] = course.Address
				courseInfo["class"] = course.RestrictClass
				courseInfo["score"] = engine.GetStudentCourseScore(student.HumanID, course.CourseID)
				if courseInfo["score"].(*int) != nil {
					courseInfo["score"] = strconv.Itoa(*courseInfo["score"].(*int))
				}

				studentInfo["schedule"] = append(studentInfo["schedule"].([]gin.H), courseInfo)
			}

			response["data"] = append(response["data"].([]gin.H), studentInfo)
		}

		response.SetCodeAndMsg(0, "已读取。")
		c.JSON(http.StatusAccepted, response)
	}
}
