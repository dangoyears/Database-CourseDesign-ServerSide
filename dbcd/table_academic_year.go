package dbcd

import (
	"log"
)

// AcademicYear 表“AcademicYear”的模型
type AcademicYear struct {
	AcademicYear int
}

// CreateAcademicYear 创建年份为year的学年。
func (engine *Engine) CreateAcademicYear(year int) {
	query := `insert into "AcademicYear" values (:1)`
	_, err := engine.db.Exec(query, year)
	if err != nil {
		log.Println(err)
	}
}

// DeleteAcademicYear 删除年份为year的学年。
func (engine *Engine) DeleteAcademicYear(yaer int) {
	query := `delete from "AcademicYear" where "AcademicYear"."AcademicYear"=:1`
	_, err := engine.db.Exec(query, yaer)
	if err != nil {
		log.Println(err)
	}
}

// TestAcademicYear 测试能否向表中插入数据
func (engine *Engine) TestAcademicYear() {
	const testYear = 3030

	log.Println("Testing table AcademicYear.")
	engine.CreateAcademicYear(testYear)
	engine.DeleteAcademicYear(testYear)
}
