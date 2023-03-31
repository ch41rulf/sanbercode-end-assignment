package repository

import (
	"database/sql"
	"end-assignment/structs"
)

func GetSubtask(db *sql.DB) (err error, results []structs.Subtask) {
	sql := "SELECT * FROM subtask"

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var subtask = structs.Subtask{}

		err = rows.Scan(&subtask)
		if err != nil {
			panic(err)
		}

		results = append(results, subtask)
	}
	return

}

func InsertsSubtask(db *sql.DB, subtask structs.Subtask) (err error) {

	sql := `INSERT INTO subtasks (subtask_id, title, description, due_date, status, main_task_id, created_at, updated_at) 
			VALUES ($1, $2, $3, $4, $5, (SELECT task_id FROM main_tasks WHERE task_id = ?), NOW(), NOW())`

	errs := db.QueryRow(sql)
	return errs.Err()

}

func UpdateSubtask(db *sql.DB, subtask structs.Subtask) (err error) {
	sql := "UPDATE subtask SET title=$1, description=$4,update_at=NOW() WHERE subtask_id=$4"

	errs := db.QueryRow(sql, subtask)

	return errs.Err()
}

func DeleteSubtask(db *sql.DB, subtask structs.Subtask) (err error) {
	sql := "DELETE FROM subtask WHERE subtask_id = $1"

	errs := db.QueryRow(sql, subtask)

	return errs.Err()
}
