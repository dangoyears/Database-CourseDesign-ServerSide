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
		Trace(err, query, year)
	}
}

// ExistAcademicYear 返回year指定的学年是否存在。
func (engine *Engine) ExistAcademicYear(year int) bool {
	query := `select count(*) from "AcademicYear" where "AcademicYear"=:1`
	result := engine.db.QueryRow(query, year)

	var count int
	err := result.Scan(&count)
	if err != nil {
		Trace(err, query, year)
	}
	return count >= 1
}

// DeleteAcademicYear 删除年份为year的学年。
func (engine *Engine) DeleteAcademicYear(year int) {
	query := `delete from "AcademicYear" where "AcademicYear"=:1`
	_, err := engine.db.Exec(query, year)
	if err != nil {
		Trace(err, query, year)
	}
}

// TestTableAcademicYear 测试能否向表中插入数据
func (engine *Engine) TestTableAcademicYear() {
	log.Println("Testing table AcademicYear.")

	// 准备测试环境。
	const testYear = 3030
	engine.DeleteAcademicYear(testYear)

	// 测试CREATE。
	engine.CreateAcademicYear(testYear)
	if !engine.ExistAcademicYear(testYear) {
		log.Panicln("Table AcademicYear test failed! Test year should exist.")
	}

	// 测试DELETE。
	engine.DeleteAcademicYear(testYear)
	if engine.ExistAcademicYear(testYear) {
		log.Panicln("Table AcademicYear test failed! Test year should NOT exist.")
	}
}
