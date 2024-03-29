package dbcd

import (
	"log"
	"time"
)

// Student 表Student的模型
type Student struct {
	HumanID        int
	ClassID        int
	StudentNumber  int
	AdmissionDate  time.Time
	GraduationDate time.Time
	StudentDegree  string
	YearOfSchool   int
	Status         string
}

// CreateStudent 创建学生。
func (engine *Engine) CreateStudent(human Human, collegeName, specialtyName string, grade, classCode, studentNumber int) {
	humanID := engine.CreateHuman(human)

	if !engine.ExistClass(collegeName, specialtyName, grade, classCode) {
		engine.CreateClass(collegeName, specialtyName, grade, classCode)
	}
	classID := engine.GetClassBySpecialtyNameGradeAndCode(specialtyName, grade, classCode).ClassID

	query := `insert into "Student" ("HumanID", "ClassID", "StudentNumber") values (:1, :2, :3)`
	_, err := engine.db.Exec(query, humanID, classID, studentNumber)
	if err != nil {
		log.Println(query, studentNumber, err)
	}
}

// ExistStudent 返回指定studentNumber的学生是否存在。
func (engine *Engine) ExistStudent(studentNumber int) bool {
	return engine.GetStudentByStudentNubmer(studentNumber) != nil
}

// GetStudentByStudentNubmer 返回对应StudentNumber的学生信息。
func (engine *Engine) GetStudentByStudentNubmer(studentNumber int) *Student {
	query := `select "HumanID", "ClassID", "StudentNumber" from "Student" where "StudentNumber"=:1`

	var student Student
	row := engine.db.QueryRow(query, studentNumber)
	if err := row.Scan(&student.HumanID, &student.ClassID, &student.StudentNumber); err != nil {
		return nil
	}
	return &student
}

// GetStudentsByClassID 获取指定ClassID的学生
func (engine *Engine) GetStudentsByClassID(classID int) []Student {
	query := `select "StudentNumber" from "Student" where "ClassID"=:1`
	rows, err := engine.db.Query(query, classID)
	if err != nil {
		Trace(err, query, classID)
	}
	defer rows.Close()

	var studentNumbers []int
	for rows.Next() {
		var number int
		if err := rows.Scan(&number); err != nil {
			Trace(err)
		}
		studentNumbers = append(studentNumbers, number)
	}

	var students []Student
	for _, number := range studentNumbers {
		student := engine.GetStudentByStudentNubmer(number)
		if student != nil {
			students = append(students, *student)
		}
	}
	return students
}

// UpdateStudent 更新指定学号studentNumber学生的信息。
// 只能更新学生本身的属性，不能处理转专业等情形。
// 如果需要转专业等，使用UpdateStudentAsInfo。
func (engine *Engine) UpdateStudent(studentNumber int, info StudentInfo) {
	query := `update "Student" set "AdmissionDate"=:1, "GraduationDate"=:2, "StudentDegree"=:3, "YearOfSchool"=:4, "Status"=:5 where "StudentNumber"=:6`
	_, err := engine.db.Exec(query, info.AdmissionDate, info.GraduationDate, info.StudentDegree, info.YearOfSchool, info.Status, studentNumber)
	if err != nil {
		log.Println(query, studentNumber, err)
	}
}

// DeleteStudentByID 删除指定HumanID的学生。
func (engine *Engine) DeleteStudentByID(id int) {
	query := `delete from "Student" where "HumanID"=:1`
	_, err := engine.db.Exec(query, id)
	if err != nil {
		Trace(err, query, id)
	}

	engine.DeleteHumanByID(id)
}

// DeleteStudentByStudentNubmer 删除指定ID的学生。
func (engine *Engine) DeleteStudentByStudentNubmer(studentNumber int) {
	student := engine.GetStudentByStudentNubmer(studentNumber)
	if student == nil {
		return
	}

	query := `delete from "Student" where "StudentNumber"=:1`
	_, err := engine.db.Exec(query, studentNumber)
	if err != nil {
		Trace(err, query, studentNumber)
	}

	engine.DeleteHumanByID(student.HumanID)
}

// DeleteStudentsByClassID 删除指定ClassID的学生。
func (engine *Engine) DeleteStudentsByClassID(classID int) {
	students := engine.GetStudentsByClassID(classID)

	for _, student := range students {
		engine.DeleteStudentByStudentNubmer(student.StudentNumber)
	}
}

// TestTableStudent 测试表Student。
func (engine *Engine) TestTableStudent() {
	log.Println("Testing table Student.")

	// 准备测试环境。
	const (
		testName          = "（测试学生）"
		testIdentity      = "123456789012345678"
		testCollegeName   = "（测试学院）"
		testSpecialtyName = "（测试专业）"
		testGrade         = 10
		testClassCode     = 24
		testStudentNumber = 1706300000
	)
	var human = Human{
		Name:     testName,
		Identity: testIdentity,
	}
	engine.DeleteStudentByStudentNubmer(testStudentNumber)

	// 测试CREATE。
	engine.CreateStudent(human, testCollegeName, testSpecialtyName, testGrade, testClassCode, testStudentNumber)
	if !engine.ExistStudent(testStudentNumber) {
		log.Panicln("Table Student test failed: student should exist.")
	}

	// 测试DELETE。
	engine.DeleteStudentByStudentNubmer(testStudentNumber)
	if engine.ExistStudent(testStudentNumber) {
		log.Panicln("Table Student test failed: student should NOT exist.")
	}
}
