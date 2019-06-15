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

// CreateAndGetSpecialty 创建并返回由参数指定的Specialty结构的指针。
// 若Specialty已经存在，则返回现有结构的指针。
func (engine *Engine) CreateAndGetSpecialty(collegeName, specialtyName string) {
	college := engine.CreateAndGetCollege(collegeName)
}

// TestSpecialty 测试表Specialty。
func (engine *Engine) TestSpecialty() {
	log.Println("Test table Specialty.")
	
	const (
		testCollege = "如果在看见数据表中看见此学院，后端的数据库测试可能没有成功。"
		testSpecialty = "如果在看见数据表中看见此专业，后端的数据库测试可能没有成功。"
	)
}
