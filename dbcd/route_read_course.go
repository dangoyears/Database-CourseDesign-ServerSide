package dbcd

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetReadCourseEndpoint 返回“/read/course”处的路由。
func (engine *Engine) GetReadCourseEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		var response = NewRouterResponse()

		response["data"] = []gin.H{}

		courses := engine.GetCourse()

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

			courseInfo["students"] = []gin.H{}
			studentHumanIDs := engine.GetStudentHumanIDsByCourseID(course.CourseID)
			for _, id := range studentHumanIDs {
				info := engine.GetStudentInfoByStudentHumanID(id)

				studentInfo := make(gin.H)
				studentInfo["college"] = info.CollegeName
				studentInfo["specialty"] = info.SpecialtyName
				studentInfo["grade"] = strconv.Itoa(info.Grade)
				studentInfo["class"] = strconv.Itoa(info.ClassCode)
				studentInfo["name"] = info.Name
				studentInfo["studentId"] = fmt.Sprintf("%10d", info.StudentNumber)
				studentInfo["sex"] = info.Sex

				courseInfo["students"] = append(courseInfo["students"].([]gin.H), studentInfo)
			}

			response["data"] = append(response["data"].([]gin.H), courseInfo)
		}

		response.SetCodeAndMsg(0, "已读取。")
		c.JSON(http.StatusAccepted, response)
	}
}
