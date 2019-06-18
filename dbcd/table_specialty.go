package dbcd

import (
	"log"
)

// Specialty 是表“Specialty”的模型
type Specialty struct {
	SpecialtyID   int
	CollegeID     int
	SpecialtyName string
}

// CreateSpecialty 在指定学院下创建专业。
func (engine *Engine) CreateSpecialty(collegeName, specialtyName string) {
	query := `insert into "Specialty" ("CollegeID", "SpecialtyName") values (:1, :2)`

	if !engine.ExistCollege(collegeName) {
		engine.CreateCollege(collegeName)
	}
	college := engine.GetCollegeByName(collegeName)

	_, err := engine.db.Exec(query, college.CollegeID, specialtyName)
	if err != nil {
		log.Println(query, collegeName, specialtyName, err)
	}
}

// ExistSpecialty 返回具有name名称的专业是否存在。
func (engine *Engine) ExistSpecialty(name string) bool {
	query := `select count(*) from "Specialty" where "SpecialtyName"=:1`
	result := engine.db.QueryRow(query, name)

	var count int
	if err := result.Scan(&count); err != nil {
		log.Println(name, err)
	}
	return count >= 1
}

// GetSpecialtyByName 返回由名称指定的专业。
func (engine *Engine) GetSpecialtyByName(name string) *Specialty {
	query := `select "SpecialtyID", "CollegeID", "SpecialtyName" from "Specialty" where "SpecialtyName"=:1`
	result := engine.db.QueryRow(query, name)

	var specialty Specialty
	if err := result.Scan(&specialty.SpecialtyID, &specialty.CollegeID, &specialty.SpecialtyName); err != nil {
		return nil
	}
	return &specialty
}

// DeleteSpecialtyByName 删除指定名称name的专业。
func (engine *Engine) DeleteSpecialtyByName(name string) {
	specialty := engine.GetSpecialtyByName(name)
	if specialty == nil {
		return
	}
	id := specialty.SpecialtyID

	queryDeleteChildRecords := `delete from "Class" where "SpecialtyID"=:1`
	if _, err := engine.db.Exec(queryDeleteChildRecords, id); err != nil {
		log.Println(queryDeleteChildRecords, id, err)
	}

	query := `delete from "Specialty" where "SpecialtyID"=:1`
	_, err := engine.db.Exec(query, id)
	if err != nil {
		log.Println(query, id, err)
	}
}

// TestTableSpecialty 测试表Specialty。
func (engine *Engine) TestTableSpecialty() {
	log.Println("Testing table Specialty.")

	// 准备测试环境。
	const (
		testCollegeName   = "（测试学院）"
		testSpecialtyName = "（测试专业）"
	)
	engine.DeleteSpecialtyByName(testSpecialtyName)
	engine.DeleteCollegeByName(testCollegeName)

	// 测试CREATE。
	engine.CreateSpecialty(testCollegeName, testSpecialtyName)
	if !engine.ExistSpecialty(testSpecialtyName) {
		log.Panicln("Table Specialty test failed! Specialty with testSpecialtyName should exist.")
	}

	// 测试READ。
	specialty := engine.GetSpecialtyByName(testSpecialtyName)
	if specialty == nil {
		log.Panicln("Table Specialty test failed: specialty should NOT be nil!")
	}
	if specialty.CollegeID != engine.GetCollegeByName(testCollegeName).CollegeID {
		log.Panicln("Table Specialty test failed! College ID should be same.")
	}

	// 测试DELETE。
	engine.DeleteSpecialtyByName(testSpecialtyName)
	engine.DeleteCollegeByName(testCollegeName)
	if engine.ExistSpecialty(testSpecialtyName) {
		log.Panicln("Table Specialty test failed: Specialty with testSpecialtyName should NOT exist!")
	}
}
