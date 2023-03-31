package repository

import (
	"database/sql"
	"end-assignment/structs"
)

func GetMainTask(db *sql.DB) (err error, results []structs.MainTask) {
	sql := "SELECT * FROM main_tasks"

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var maintask = structs.MainTask{}

		err = rows.Scan(&maintask)
		if err != nil {
			panic(err)
		}

		results = append(results, maintask)
	}
	return

}

func InsertsMainTask(db *sql.DB, maintask structs.MainTask) (err error) {

	sql := `INSERT INTO main_tasks(title, description, due_date, status, manager_id, created_at) 
			VALUES ($1, $2, $3, $4,(SELECT user_id FROM users WHERE users_id = $5), NOW(), NOW())`
	errs := db.QueryRow(sql)
	return errs.Err()

}

func UpdateMainTask(db *sql.DB, maintask structs.MainTask) (err error) {
	sql := "UPDATE main_tasks SET title=$1, description=$2, due_date=$3, status=$4,updated_at=NOW() WHERE task_id=$6"

	errs := db.QueryRow(sql, maintask)

	return errs.Err()
}

func DeleteMaintask(db *sql.DB, maintask structs.MainTask) (err error) {
	sql := "DELETE FROM main_tasks WHERE task_id = $1"

	errs := db.QueryRow(sql, maintask)

	return errs.Err()
}
