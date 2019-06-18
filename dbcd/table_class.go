package dbcd

import "log"

// Class 是表“Class”的抽象
type Class struct {
	ClassID              int
	SpecialtyID          int
	MasterTeacherHumanID *int
	Grade                int
	ClassCode            int
}

// CreateClass 根据专业名specialty、届别grade和班别code来创建班级。
func (engine *Engine) CreateClass(collegeName, specialtyName string, grade, code int) {
	if !engine.ExistSpecialty(specialtyName) {
		engine.CreateSpecialty(collegeName, specialtyName)
	}
	specialtyID := engine.GetSpecialtyByName(specialtyName).SpecialtyID

	query := `insert into "Class" ("SpecialtyID", "MasterTeacherHumanID", "Grade", "ClassCode")
values (:1, null, :2, :3)`
	_, err := engine.db.Exec(query, specialtyID, grade, code)
	if err != nil {
		Trace(err, query, specialtyID, grade, code)
	}
}

// ExistClass 返回指定专业、届别和班别的班级是否存在
func (engine *Engine) ExistClass(collegeName, specialtyName string, grade, code int) bool {
	specialty := engine.GetSpecialtyByName(specialtyName)
	if specialty == nil {
		return false
	}
	specialtyID := specialty.SpecialtyID

	query := `select count(*) from "Class" where "SpecialtyID"=:1 and "Grade"=:2 and "ClassCode"=:3`
	result := engine.db.QueryRow(query, specialtyID, grade, code)

	var count int
	if err := result.Scan(&count); err != nil {
		Trace(err, query, collegeName, specialtyName, grade, code)
	}
	return count >= 1
}

// GetClassBySpecialtyNameGradeAndCode 根据班级的专业名称、届别和班别取得班级信息。
func (engine *Engine) GetClassBySpecialtyNameGradeAndCode(specailtyName string, grade, code int) *Class {
	specialty := engine.GetSpecialtyByName(specailtyName)
	if specialty == nil {
		return nil
	}
	specialtyID := specialty.SpecialtyID

	query := `select "ClassID", "SpecialtyID", "MasterTeacherHumanID", "Grade", "ClassCode" from "Class"
where "SpecialtyID"=:1 and "Grade"=:2 and "ClassCode"=:3`
	result := engine.db.QueryRow(query, specialtyID, grade, code)

	var class Class
	if err := result.Scan(&class.ClassID, &class.SpecialtyID, &class.MasterTeacherHumanID, &class.Grade, &class.ClassCode); err != nil {
		return nil
	}
	return &class
}

// DeleteClassByID 删除由CollegeID指定的班级。
func (engine *Engine) DeleteClassByID(id int) {
	engine.DeleteStudentsByClassID(id)

	query := `delete from "Class" where "ClassID"=:1`
	_, err := engine.db.Exec(query, id)
	if err != nil {
		Trace(query, id, err)
	}
}

// DeleteClassBySpecialtyNameGradeAndCode 删除由专业、届别和班别指定的班级。
func (engine *Engine) DeleteClassBySpecialtyNameGradeAndCode(specialtyName string, grade, code int) {
	class := engine.GetClassBySpecialtyNameGradeAndCode(specialtyName, grade, code)
	if class == nil {
		return
	}

	engine.DeleteClassByID(class.ClassID)
}

// TestTableClass 测试Class表。
func (engine *Engine) TestTableClass() {
	log.Println("Testint table Class.")

	// 准备测试环境。
	const (
		testCollegeName   = "（测试学院）"
		testSpecialtyName = "（测试专业）"
		testGrade         = 10
		testClassCode     = 24
	)
	engine.DeleteClassBySpecialtyNameGradeAndCode(testSpecialtyName, testGrade, testClassCode)

	// 测试CREATE。
	engine.CreateClass(testCollegeName, testSpecialtyName, testGrade, testClassCode)
	if !engine.ExistClass(testCollegeName, testSpecialtyName, testGrade, testClassCode) {
		log.Panicln("Table Class test failed: testClass should exist.")
	}

	// 测试READ。
	class := engine.GetClassBySpecialtyNameGradeAndCode(testSpecialtyName, testGrade, testClassCode)
	if class.MasterTeacherHumanID != nil {
		log.Panicln("Table Class test failed: MasterTeacherHumanID should be nil.")
	}
	if class.Grade != testGrade || class.ClassCode != testClassCode {
		log.Panicln("Table Class test failed: Grade or ClassCode miss match.")
	}

	// 测试DELETE。
	engine.DeleteClassBySpecialtyNameGradeAndCode(testSpecialtyName, testGrade, testClassCode)
	if engine.ExistClass(testCollegeName, testCollegeName, testGrade, testClassCode) {
		log.Panicln("Table Class test failed: testClass should NOT exist.")
	}

	// 清理测试环境。
	engine.DeleteCollegeByName(testCollegeName)
}
