package dbcd

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// GetReadTeacherOneEndpoint 返回“/read/teacher/one”处的路由。
func (engine *Engine) GetReadTeacherOneEndpoint() gin.HandlerFunc {

	type readTeacherOneEndpointParam struct {
		TeacherNumber string `json:"jobId" form:"jobId" binding:"required"`
	}

	return func(c *gin.Context) {
		resumeRequestBody(c)

		var param readTeacherOneEndpointParam
		var response = NewRouterResponse()

		if c.ShouldBind(&param) == nil {
			teacherNumber, err := strconv.Atoi(param.TeacherNumber)
			if err != nil {
				response.SetCodeAndMsg(-1, "jobId必须可以转换为整数。")
				c.JSON(http.StatusOK, response)
				return
			}

			response["data"] = []gin.H{}
			teachers := engine.GetTeacherInfo()
			for _, teacher := range teachers {
				if teacher.TeacherNumber != teacherNumber {
					continue
				}

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

					teacherInfo["schedule"] = append(teacherInfo["schedule"].([]gin.H), courseInfo)
				}

				response["data"] = append(response["data"].([]gin.H), teacherInfo)
			}

			if len(response["data"].([]gin.H)) >= 1 {
				response["data"] = response["data"].([]gin.H)[0]
			} else {
				response["data"] = nil
			}

			response.SetCodeAndMsg(0, "已读取。")
			c.JSON(http.StatusAccepted, response)
			return
		}

		response.SetCodeAndMsg(-1, "参数不足，必须提供jobId参数。")
		c.JSON(http.StatusOK, response)
	}
}
