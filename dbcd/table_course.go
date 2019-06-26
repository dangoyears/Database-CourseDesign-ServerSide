package dbcd

import "strconv"

// Course 是Course表的抽象
type Course struct {
	CourseID           int
	LeadTeacherHumanID *int
	CourseName         string
	CourseNumber       int
	Credits            int
	CourseProperty     int
	Accommodate        int
	Time               string
	Address            string
	RestrictClass      string
}

// CreateCourse 按照指定的信息创建课程。
func (engine *Engine) CreateCourse(courseName string, courseNumber, credits, courseProperty, accommodate int, time, address, restrictClass string) {
	query := `insert into "Course" ("CourseName", "CourseNumber", "Credits", "CourseProperty", "Accommodate", "Time", "Address", "RestrictClass")
values (:1, :2, :3, :4, :5, :6, :7, :8)`
	_, err := engine.db.Exec(query, courseName, courseNumber, credits, courseProperty, accommodate, time, address, restrictClass)
	if err != nil {
		Trace(err, query, courseName, courseNumber, credits, courseProperty, accommodate, time, address, restrictClass)
	}
}

// GetCourse 获取所有课程。
func (engine *Engine) GetCourse() []Course {
	query := `select "CourseID", "LeadTeacherHumanID", "CourseName", "CourseNumber", "Credits", "CourseProperty", "Accommodate", "Time", "Address", "RestrictClass" from "Course"`
	rows, err := engine.db.Query(query)
	if err != nil {
		Trace(err, query)
	}
	defer rows.Close()

	var courses []Course
	for rows.Next() {
		var course Course
		if err := rows.Scan(&course.CourseID, &course.LeadTeacherHumanID, &course.CourseName, &course.CourseNumber,
			&course.Credits, &course.CourseProperty, &course.Accommodate, &course.Time, &course.Address, &course.RestrictClass); err != nil {
			Trace(query, err)
		}
		courses = append(courses, course)
	}

	return courses
}

// GetCourseByCourseNumber 获取指定courseNumber的课程。
func (engine *Engine) GetCourseByCourseNumber(courseNumber int) *Course {
	query := `select "CourseID", "LeadTeacherHumanID", "CourseName", "CourseNumber", "Credits", "CourseProperty", "Accommodate", "Time", "Address", "RestrictClass" 
from "Course" where "CourseNumber"=:1`
	row := engine.db.QueryRow(query, courseNumber)

	var course Course
	if err := row.Scan(&course.CourseID, &course.LeadTeacherHumanID, &course.CourseName, &course.CourseNumber,
		&course.Credits, &course.CourseProperty, &course.Accommodate, &course.Time, &course.Address, &course.RestrictClass); err != nil {
		return nil
	}
	return &course
}

// UpdateCourseByCourseName 按照指定的信息创建课程。
func (engine *Engine) UpdateCourseByCourseName(courseNumber int, courseName string, credits, courseProperty, accommodate int, time, address, restrictClass string) {
	query := `update "Course" set "CourseName"=:1, "Credits"=:2, "CourseProperty"=:3, "Accommodate"=:4, "Time"=:5, "Address"=:6, "RestrictClass"=:7 where "CourseNumber"=:8`
	_, err := engine.db.Exec(query, courseName, credits, courseProperty, accommodate, time, address, restrictClass, courseNumber)
	if err != nil {
		Trace(err, query, courseName, courseNumber, credits, courseProperty, accommodate, time, address, restrictClass)
	}
}

// UpdateCourseLeadTeacherID 更新指定课程的课程负责人，可以设置为空。
func (engine *Engine) UpdateCourseLeadTeacherID(courseNumber int, leadTeacherID *int) {
	query := `update "Course" set "LeadTeacherHumanID"=:1 where "CourseNumber"=:2`
	_, err := engine.db.Exec(query, leadTeacherID, courseNumber)

	if err != nil {
		Trace(err, query, leadTeacherID, courseNumber)
	}
}

// AddTeacherByCourseNumberAndTeacherHumanID 添加授课教师。
func (engine *Engine) AddTeacherByCourseNumberAndTeacherHumanID(courseNumber, teacherHumanID int) {
	course := engine.GetCourseByCourseNumber(courseNumber)
	if course == nil {
		Trace("未找到courseNumber对应的课程：" + strconv.Itoa(courseNumber))
	}

	query := `insert into "TeacherTeachsCourse" ("CourseID", "TeacherHumanID") values (:1, :2)`
	_, err := engine.db.Exec(query, course.CourseID, teacherHumanID)
	if err != nil {
		Trace(err, query)
	}
}

// AddCourseForClass 为班级添加课程。
func (engine *Engine) AddCourseForClass(courseNumber, classID int) {
	course := engine.GetCourseByCourseNumber(courseNumber)
	if course == nil {
		Trace("未找到courseNumber对应的课程：" + strconv.Itoa(courseNumber))
	}

	students := engine.GetStudentsByClassID(classID)
	for _, student := range students {
		query := `insert into "StudentAttendsCourse" ("CourseID", "StudentHumanID") values (:1, :2)`
		_, err := engine.db.Exec(query, course.CourseID, student.HumanID)
		if err != nil {
			Trace(err, query, course.CourseID, student.HumanID)
		}
	}
}

// RemoveCourseForClass 为班级删除课程。
func (engine *Engine) RemoveCourseForClass(courseNumber, classID int) {
	course := engine.GetCourseByCourseNumber(courseNumber)
	if course == nil {
		Trace("未找到courseNumber对应的课程：" + strconv.Itoa(courseNumber))
	}

	students := engine.GetStudentsByClassID(classID)
	for _, student := range students {
		query := `delete from "StudentAttendsCourse" where "CourseID"=:1 and "StudentHumanID"=:2`
		_, err := engine.db.Exec(query, course.CourseID, student.HumanID)
		if err != nil {
			Trace(err, query, course.CourseID, student.HumanID)
		}
	}
}

// RemoveAllTeacher 移除除课程负责人以外的所有教师。
func (engine *Engine) RemoveAllTeacher(courseNumber int) {
	course := engine.GetCourseByCourseNumber(courseNumber)
	if course == nil {
		Trace("未找到courseNumber对应的课程：" + strconv.Itoa(courseNumber))
	}

	query := `delete from "TeacherTeachsCourse" where "CourseID"=:1`
	_, err := engine.db.Exec(query, course.CourseID)
	if err != nil {
		Trace(err, query)
	}
}

// DeleteCourseByCourseNumber 删除指定courseNubmer的课程。
func (engine *Engine) DeleteCourseByCourseNumber(courseNumber int) {
	query := `delete from "Course" where "CourseNumber"=:1`
	_, err := engine.db.Exec(query, courseNumber)

	if err != nil {
		Trace(err, query)
	}
}

// TestTableCourse 用于测试表的功能。
func (engine *Engine) TestTableCourse() {
	// 未添加单元测试
}
