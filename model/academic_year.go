package model

import (
	"database/sql"
	"log"
)

type AcademicYear struct {
	AcademicYear int
}

func testInsertIntoAcademicYear(db *sql.DB) {
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
