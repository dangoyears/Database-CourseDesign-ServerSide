package dbcd

import "log"

// Teacher 是表Teacher的抽象
type Teacher struct {
	HumanID          int
	CollegeID        int
	TeacherNumber    int
	GraduationSchool string
	Position         string
	TeacherDegree    string
}

// CreateTeacher 创建教师。
func (engine *Engine) CreateTeacher(human Human, collegeName string, teacherNumber int, graduationSchool, position, teacherDegree string) {
	humanID := engine.CreateHuman(human)

	if !engine.CollegeExists(collegeName) {
		engine.CreateCollege(collegeName)
	}
	collegeID := engine.GetCollegeByName(collegeName).CollegeID

	query := `insert into "Teacher" ("HumanID", "CollegeID", "TeacherNumber", "GraduationSchool", "Position", "TeacherDegree")
values (:1, :2, :3, :4, :5, :6)`
	_, err := engine.db.Exec(query, humanID, collegeID, teacherNumber, graduationSchool, position, teacherDegree)
	if err != nil {
		log.Print(query, teacherNumber, err)
	}
}

// TeacherExists 返回指定teachernNumber的教师是否存在。
func (engine *Engine) TeacherExists(teacherNumber int) bool {
	return engine.GetTeacherByTeacherNumber(teacherNumber) != nil
}

// GetTeacherByTeacherNumber 返回指定teacherNumber的教师。
func (engine *Engine) GetTeacherByTeacherNumber(teacherNumber int) *Teacher {
	query := `select "HumanID", "CollegeID", "TeacherNumber", "GraduationSchool", "Position", "TeacherDegree" from "Teacher" where "TeacherNumber"=:1`

	var teacher Teacher
	row := engine.db.QueryRow(query, teacherNumber)
	if err := row.Scan(&teacher.HumanID, &teacher.CollegeID, &teacher.TeacherNumber, &teacher.GraduationSchool, &teacher.Position, &teacher.TeacherDegree); err != nil {
		return nil
	}
	return &teacher
}

// DeleteTeacherByTeacherNumber 删除指定teacherNumber的教师。
func (engine *Engine) DeleteTeacherByTeacherNumber(teacherNumber int) {
	teacher := engine.GetTeacherByTeacherNumber(teacherNumber)
	if teacher == nil {
		return
	}

	queryDeleteTeacher := `delete from "Teacher" where "TeacherNumber"=:1`
	_, err := engine.db.Exec(queryDeleteTeacher, teacherNumber)
	if err != nil {
		log.Println(queryDeleteTeacher, teacherNumber, err)
	}

	engine.DeleteHumanByID(teacher.HumanID)
}

// TestTableTeacher 测试表Teacher。
func (engine *Engine) TestTableTeacher() {
	log.Println("Testing table Teacher.")

	const (
		testName             = "（测试名称）"
		testIdentity         = "123456789012345678"
		testCollegeName      = "（测试学院）"
		testTeacherName      = "（测试教师）"
		testTeacherNumber    = 1234567890
		testGraduationSchool = "（测试毕业院校）"
		testPosition         = "教务办主任"
		testTeacherDegree    = "博士后"
	)

	engine.DeleteTeacherByTeacherNumber(testTeacherNumber)

	var human = Human{
		Name:     testName,
		Identity: testIdentity,
	}
	engine.CreateTeacher(human, testCollegeName, testTeacherNumber, testGraduationSchool, testPosition, testTeacherDegree)
	if !engine.TeacherExists(testTeacherNumber) {
		log.Panicln("Table Teacher test failed: teacher should exist.")
	}
	engine.DeleteTeacherByTeacherNumber(testTeacherNumber)
	if engine.TeacherExists(testTeacherNumber) {
		log.Panicln("Table Teacher test failed: teacher should NOT exist.")
	}
}
