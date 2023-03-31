package repository

import (
	"database/sql"
	"end-assignment/structs"
)

func GetAssgnmnt(db *sql.DB) (err error, results []structs.Assignment) {
	sql := "SELECT * FROM assignmentss"

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var assgnmnt = structs.Assignment{}

		err = rows.Scan(&assgnmnt)
		if err != nil {
			panic(err)
		}

		results = append(results, assgnmnt)
	}
	return

}

func InsertsAssgnmnt(db *sql.DB, assgnmnt structs.Assignment) (err error) {

	sql := "INSERT INTO assignmentss ()"
	errs := db.QueryRow(sql)
	return errs.Err()

}

func UpdateAssgnmnt(db *sql.DB, assgnmnt structs.Assignment) (err error) {
	sql := "UPDATE assignmentss SET "

	errs := db.QueryRow(sql, assgnmnt)

	return errs.Err()
}

func DeleteAssgnmnt(db *sql.DB, assgnmnt structs.Assignment) (err error) {
	sql := "DELETE FROM assignmentss WHERE id = $1"

	errs := db.QueryRow(sql, assgnmnt)

	return errs.Err()
}
