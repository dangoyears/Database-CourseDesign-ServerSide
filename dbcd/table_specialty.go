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

// SpecialtyExists 返回具有name名称的专业是否存在。
func (engine *Engine) SpecialtyExists(name string) bool {
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
	specailty := engine.GetSpecialtyByName(name)
	if specailty == nil {
		return
	}
	id := specailty.SpecialtyID

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

// CreateSpecialty 在指定学院下创建专业。
func (engine *Engine) CreateSpecialty(collegeName, specialtyName string) {
	query := `insert into "Specialty" ("CollegeID", "SpecialtyName") values (:1, :2)`

	if !engine.CollegeExists(collegeName) {
		engine.CreateCollege(collegeName)
	}
	college := engine.GetCollegeByName(collegeName)

	_, err := engine.db.Exec(query, college.CollegeID, specialtyName)
	if err != nil {
		log.Println(query, collegeName, specialtyName, err)
	}
}

// TestTableSpecialty 测试表Specialty。
func (engine *Engine) TestTableSpecialty() {
	log.Println("Testing table Specialty.")

	const (
		testCollegeName   = "若此学院可见，数据表测试可能没有成功。"
		testSpecialtyName = "若此专业可见，数据表测试可能没有成功。"
	)

	engine.DeleteSpecialtyByName(testSpecialtyName)
	engine.DeleteCollegeByName(testCollegeName)

	engine.CreateSpecialty(testCollegeName, testSpecialtyName)
	if !engine.SpecialtyExists(testSpecialtyName) {
		log.Panicln("Table Specialty test failed! Specialty with testSpecialtyName should exist.")
	}

	specialty := engine.GetSpecialtyByName(testSpecialtyName)
	if specialty == nil {
		log.Panicln("Table Specialty test failed: specialty should NOT be nil!")
	}
	if specialty.CollegeID != engine.GetCollegeByName(testCollegeName).CollegeID {
		log.Panicln("Table Specialty test failed! College ID should be same.")
	}

	engine.DeleteSpecialtyByName(testSpecialtyName)
	engine.DeleteCollegeByName(testCollegeName)

	if engine.SpecialtyExists(testSpecialtyName) {
		log.Panicln("Table Specialty test failed: Specailty with testSpecialtyName should NOT exist!")
	}
}
