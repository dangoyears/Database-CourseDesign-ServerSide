package dbcd

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// GetReadTeacherEndpoint 返回“/read/teacher”处的路由。
func (engine *Engine) GetReadTeacherEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		var response = NewRouterResponse()

		response["data"] = []gin.H{}

		teachers := engine.GetTeacherInfo()

		for _, teacher := range teachers {
			teacherInfo := make(gin.H)

			teacherInfo["college"] = teacher.CollegeName
			teacherInfo["name"] = teacher.Name
			teacherInfo["jobId"] = fmt.Sprintf("%010d", teacher.TeacherNumber)
			teacherInfo["sex"] = teacher.Sex
			teacherInfo["education"] = teacher.TeacherDegree
			teacherInfo["graduation"] = teacher.GraduationSchool
			teacherInfo["birthday"] = fmt.Sprintf("%d-%d-%d", teacher.Birthday.Year(), teacher.Birthday.Month(), teacher.Birthday.Day())
			teacherInfo["age"] = fmt.Sprintf("%d", time.Now().Year()-teacher.Birthday.Year())
			teacherInfo["idCard"] = teacher.Identity
			teacherInfo["position"] = teacher.Position
			teacherInfo["schedule"] = []gin.H{}

			courses := engine.GetCourseByTeacherNumber(teacher.TeacherNumber)
			for _, course := range courses {
				courseInfo := make(gin.H)

				courseInfo["name"] = course.CourseName
				courseInfo["id"] = fmt.Sprintf("%010d", course.CourseNumber)
				courseInfo["credit"] = course.Credits
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

				teacherInfo["schedule"] = append(teacherInfo["schedule"].([]gin.H), courseInfo)
			}

			response["data"] = append(response["data"].([]gin.H), teacherInfo)
		}

		response.SetCodeAndMsg(0, "已读取。")
		c.JSON(http.StatusAccepted, response)
	}
}
