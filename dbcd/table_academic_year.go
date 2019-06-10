package dbcd

import (
	"database/sql"
	"log"
)

// AcademicYear 是表“AcademicYear”的模型
type AcademicYear struct {
	AcademicYear int
}

// TestInsertIntoAcademicYear 测试能否向表中插入数据
func TestInsertIntoAcademicYear(db *sql.DB) {
	var (
		err      error
		testYear = 3030
	)

	_, err = db.Exec(`INSERT INTO "AcademicYear" ("AcademicYear") VALUES (:1)`, testYear)
	if err != nil {
		log.Println(err)
		return
	}
	_, err = db.Exec(`DELETE FROM "AcademicYear" WHERE "AcademicYear"=:1`, testYear)
	if err != nil {
		log.Println(err)
	}
}
