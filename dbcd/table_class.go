package dbcd

import "log"

// Class 是表“Class”的抽象
type Class struct {
	ClassID              int
	SpecialtyID          int
	MasterTeacherHumanID int
	Grade                int
	ClassCode            int
}

// ClassExists 返回指定专业、届别和班别的班级是否存在
func (engine *Engine) ClassExists(collegeName, specialtyName string, grade, code int) bool {
	specialty := engine.GetSpecialtyByName(specialtyName)
	if specialty == nil {
		return false
	}
	specialtyID := specialty.SpecialtyID

	query := `select count(*) from "Class" where "SpecialtyID"=:1 and "Grade"=:2 and "ClassCode"=:3`
	result := engine.db.QueryRow(query, specialtyID, grade, code)

	var count int
	if err := result.Scan(&count); err != nil {
		log.Println(query, collegeName, specialtyName, grade, code)
	}
	return count >= 1
}

// CreateClass 根据专业名specialty、届别grade和班别code来创建班级。
func (engine *Engine) CreateClass(collegeName, specialtyName string, grade, code int) {
	if !engine.SpecialtyExists(specialtyName) {
		engine.CreateSpecialty(collegeName, specialtyName)
	}
	specialtyID := engine.GetSpecialtyByName(specialtyName).SpecialtyID

	query := `insert into "Class" ("SpecialtyID", "MasterTeacherHumanID", "Grade", "ClassCode")
values (:1, null, :2, :3)`
	_, err := engine.db.Exec(query, specialtyID, grade, code)
	if err != nil {
		log.Println(query, specialtyID, grade, code, err)
	}
}

// DeleteClassBySpecialtyNameGradeAndCode 删除由专业、届别和班别指定的班级。
func (engine *Engine) DeleteClassBySpecialtyNameGradeAndCode(specialtyName string, grade, code int) {
	specialty := engine.GetSpecialtyByName(specialtyName)
	if specialty == nil {
		return
	}
	specialtyID := specialty.SpecialtyID

	query := `delete from "Class" where "SpecialtyID"=:1 and "Grade"=:2 and "ClassCode"=:3`
	_, err := engine.db.Exec(query, specialtyID, grade, code)
	if err != nil {
		log.Println(query, specialtyName, grade, code, err)
	}
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
		log.Println(query, specailtyName, grade, code, err)
		return nil
	}
	return &class
}

// TestTableClass 测试Class表。
func (engine *Engine) TestTableClass() {
	log.Println("Testint table Class.")

	const (
		testCollegeName   = "如果此学院可见，数据表测试可能没有成功。"
		testSpecialtyName = "如果此专业可见，数据表测试可能没有成功。"
		testGrade         = 10
		testClassCode     = 24
	)

	engine.DeleteClassBySpecialtyNameGradeAndCode(testSpecialtyName, testGrade, testClassCode)

	engine.CreateClass(testCollegeName, testSpecialtyName, testGrade, testClassCode)
	if !engine.ClassExists(testCollegeName, testSpecialtyName, testGrade, testClassCode) {
		log.Panicln("Table Class test failed: testClass should exist.")
	}

	engine.DeleteClassBySpecialtyNameGradeAndCode(testSpecialtyName, testGrade, testClassCode)
	if engine.ClassExists(testCollegeName, testCollegeName, testGrade, testClassCode) {
		log.Panicln("Table Class test failed: testClass should NOT exist.")
	}
}
