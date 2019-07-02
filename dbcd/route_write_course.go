package dbcd

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// GetWriteCourseEndpoint 返回“/write/course”处的路由。
func (engine *Engine) GetWriteCourseEndpoint() gin.HandlerFunc {

	type writeCourseEndpointParam struct {
		CourseName     string `json:"name" form:"name" binding:"required"`
		CourseNumber   string `json:"id" form:"id" binding:"required"`
		Credits        string `json:"credit" form:"credit" binding:"required"`
		CourseProperty string `json:"nature" form:"nature" binding:"required"`
		Accommodate    string `json:"accommodate" form:"accommodate" binding:"required"`
		Time           string `json:"time" form:"time" binding:"required"`
		Teachers       string `json:"teachers" form:"teachers" binding:"required"`
		CourseLeader   string `json:"courseLeader" form:"courseLeader" binding:"required"`
		Address        string `json:"address" form:"address" binding:"required"`
		RestrictClass  string `json:"class" form:"class" binding:"required"`
	}

	return func(c *gin.Context) {
		var param writeCourseEndpointParam
		var response = NewRouterResponse()

		resumeRequestBody(c)
		if c.ShouldBind(&param) == nil {
			courseNumber, courseNumberErr := strconv.Atoi(param.CourseNumber)
			if courseNumberErr != nil {
				response.SetCodeAndMsg(-1, "id必须可以转换为整数。")
				c.JSON(http.StatusOK, response)
				return
			}

			credits, creditsErr := strconv.Atoi(param.Credits)
			if creditsErr != nil {
				response.SetCodeAndMsg(-1, "credit必须可以转换为整数。")
				c.JSON(http.StatusOK, response)
				return
			}

			accommodate, accommodateErr := strconv.Atoi(param.Accommodate)
			if accommodateErr != nil {
				response.SetCodeAndMsg(-1, "accommodate必须可以转换为整数。")
				c.JSON(http.StatusOK, response)
				return
			}

			var courseProperty int
			switch param.CourseProperty {
			case "专业必修课":
				courseProperty = 1
			case "专业选修课":
				courseProperty = 2
			case "通识性选修课":
				courseProperty = 3
			case "体育选修课":
				courseProperty = 4
			default:
				response.SetCodeAndMsg(-1, "nature只支持专业选修课、专业必修课、通识性选修课和体育选修课。")
				c.JSON(http.StatusOK, response)
				return
			}

			// 假设教师的名字中没有英文逗号
			if len(param.Teachers) <= 4 {
				response.SetCodeAndMsg(1, "teachers不能为空数组。必须指定至少一位教师。")
				c.JSON(http.StatusOK, response)
				return
			}
			teacherNames := strings.Split(param.Teachers[1:len(param.Teachers)-1], ",")
			for i, str := range teacherNames {
				teacherNames[i] = strings.TrimSpace(teacherNames[i])
				teacherNames[i] = str[1 : len(str)-1]
			}

			if course := engine.GetCourseByCourseNumber(courseNumber); course != nil {
				response.SetCodeAndMsg(0, "课程将被更新。课程成绩将被重置。")

				engine.RemoveAllTeacher(courseNumber)
				oldCourse := engine.GetCourseByCourseNumber(courseNumber)
				classesToBeDeleted := strings.Split(oldCourse.RestrictClass[1:len(oldCourse.RestrictClass)-1], ",")
				for i, class := range classesToBeDeleted {
					class = strings.TrimSpace(class)
					if len(class) <= 2 {
						continue
					}

					classesToBeDeleted[i] = class[1 : len(class)-1]
					group := strings.Split(classesToBeDeleted[i], "-")

					specailtyName := group[1]
					grade, _ := strconv.Atoi(group[2][:2])
					classCode, _ := strconv.Atoi(group[2][2:])

					class := engine.GetClassBySpecialtyNameGradeAndCode(specailtyName, grade, classCode)
					if class != nil {
						engine.RemoveCourseForClass(courseNumber, class.ClassID)
					}
				}

				engine.UpdateCourseByCourseName(courseNumber, param.CourseName, credits, courseProperty, accommodate, param.Time, param.Address, param.RestrictClass)
				for _, name := range teacherNames {
					teacherInfo := engine.GetTeacherInfoByTeacherName(name)
					if teacherInfo != nil {
						engine.AddTeacherByCourseNumberAndTeacherHumanID(courseNumber, teacherInfo.HumanID)
					}
				}
				leadTeacher := engine.GetTeacherInfoByTeacherName(param.CourseLeader)
				if leadTeacher != nil {
					engine.UpdateCourseLeadTeacherID(courseNumber, &leadTeacher.HumanID)
					engine.AddTeacherByCourseNumberAndTeacherHumanID(courseNumber, leadTeacher.HumanID)
				}
				classes := strings.Split(param.RestrictClass[1:len(param.RestrictClass)-1], ",")
				for i, class := range classes {
					class = strings.TrimSpace(class)
					if len(class) <= 2 {
						continue
					}

					classes[i] = class[1 : len(class)-1]
					group := strings.Split(classes[i], "-")

					specailtyName := group[1]
					grade, _ := strconv.Atoi(group[2][:2])
					classCode, _ := strconv.Atoi(group[2][2:])

					class := engine.GetClassBySpecialtyNameGradeAndCode(specailtyName, grade, classCode)
					if class != nil {
						engine.AddCourseForClass(courseNumber, class.ClassID)
					}
				}

				c.JSON(http.StatusOK, response)
				return
			}

			engine.CreateCourse(param.CourseName, courseNumber, credits, courseProperty, accommodate, param.Time, param.Address, param.RestrictClass)
			for _, name := range teacherNames {
				teacherInfo := engine.GetTeacherInfoByTeacherName(name)
				if teacherInfo != nil {
					engine.AddTeacherByCourseNumberAndTeacherHumanID(courseNumber, teacherInfo.HumanID)
				}
			}
			leadTeacher := engine.GetTeacherInfoByTeacherName(param.CourseLeader)
			if leadTeacher != nil {
				engine.UpdateCourseLeadTeacherID(courseNumber, &leadTeacher.HumanID)
				engine.AddTeacherByCourseNumberAndTeacherHumanID(courseNumber, leadTeacher.HumanID)
			}
			classes := strings.Split(param.RestrictClass[1:len(param.RestrictClass)-1], ",")
			for i, class := range classes {
				class = strings.TrimSpace(class)
				if len(class) <= 2 {
					continue
				}

				classes[i] = class[1 : len(class)-1]
				group := strings.Split(classes[i], "-")

				specailtyName := group[1]
				grade, _ := strconv.Atoi(group[2][:2])
				classCode, _ := strconv.Atoi(group[2][2:])

				class := engine.GetClassBySpecialtyNameGradeAndCode(specailtyName, grade, classCode)
				if class != nil {
					engine.AddCourseForClass(courseNumber, class.ClassID)
				}
			}

			response.SetCodeAndMsg(0, "课程将被创建。")
			c.JSON(http.StatusOK, response)
			return
		}
		response.SetCodeAndMsg(-1, "参数不足。")
		c.JSON(http.StatusOK, response)
	}
}
