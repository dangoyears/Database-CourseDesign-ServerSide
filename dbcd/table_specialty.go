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
	query := `select count(*) from "Specialty" where "SepcialtyName"=:1`
	result := engine.db.QueryRow(query, name)

	var count int
	if err := result.Scan(&count); err != nil {
		log.Println(name, err)
	}
	return count >= 1
}

// GetSpecialtyByName 返回由名称指定的专业。
func (engine *Engine) GetSpecialtyByName(name string) *Specialty {
	query := `select ("SepcialtyID", "CollegeID", "SepcailtyName") from "Sepcailty"
where "SepcailtyName"=:1`
	result := engine.db.QueryRow(query)

	var specialty Specialty
	if err := result.Scan(&specialty.SpecialtyID, &specialty.CollegeID, &specialty.SpecialtyName); err != nil {
		return nil
	}
	return &specialty
}

// DeleteSpecialtyByName 删除指定名称name的专业。
func (engine *Engine) DeleteSpecialtyByName(name string) {
	query := `delete from "Specialty" where "SepcailtyName"=:1`
	_, err := engine.db.Exec(query, name)
	if err != nil {
		log.Println(name, err)
	}
}

// CreateSpecialty 在指定学院下创建专业。
func (engine *Engine) CreateSpecialty(collegeName, specialtyName string) {
	log.Panicln("NOT FINISH!")
}

// TestTableSpecialty 测试表Specialty。
func (engine *Engine) TestTableSpecialty() {
	log.Println("Testing table Specialty.")

	const (
		testCollegeName   = "若此学院可见，数据表测试可能没有成功。"
		testSpecialtyName = "若此专业可见，数据表测试可能没有成功。"
	)

	// engine.DeleteSpecialtyByName(testSpecialtyName)
	// specailty := engine.GetSpecailtyByName(testSpecialtyName)
}
