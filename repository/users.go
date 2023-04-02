package repository

import (
	"database/sql"
	"end-assignment/structs"
)

func GetUsers(db *sql.DB) (err error, results []structs.Users) {
	sql := "SELECT * FROM users"

	rows, err := db.Query(sql)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var users = structs.Users{}

		err = rows.Scan(&users.UserId, &users.Username, &users.Email, &users.PhoneNumber)
		if err != nil {
			panic(err)
		}

		results = append(results, users)
	}
	return

}

func InsertsUsers(db *sql.DB, users structs.Users) error {
	sql := "INSERT INTO users (user_id, username, password, email, phone_number, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, NOW(), NOW())"
	userId := int64(len(users.Username) + len(users.Email))
	_, err := db.Exec(sql, userId, users.Username, users.Password, users.Email, users.PhoneNumber)
	return err
}

func UpdateUsers(db *sql.DB, users structs.Users) (err error) {
	sql := "UPDATE users SET password=$1, email=$2, phone_number=$3, updated_at =$4 WHERE user_id=$5"

	errs := db.QueryRow(sql, users)

	return errs.Err()
}

func DeleteUsers(db *sql.DB, users structs.Users) (err error) {
	sql := "DELETE FROM users WHERE user_id = $1"

	errs := db.QueryRow(sql, users)

	return errs.Err()
}
