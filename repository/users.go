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

		err = rows.Scan(&users.UserId, &users.Username, &users.Password, &users.Email, &users.PhoneNumber, &users.CreatedAt, &users.UpdatedAt)
		if err != nil {
			panic(err)
		}

		results = append(results, users)
	}
	return

}

func InsertsUsers(db *sql.DB, users structs.Users) error {
	sql := "INSERT INTO users (user_id, username, password, email, phone_number, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, NOW() AT TIME ZONE 'Asia/Jakarta', NOW() AT TIME ZONE 'Asia/Jakarta')"
	userId := int64(len(users.Username) + len(users.Email))
	_, err := db.Exec(sql, userId, users.Username, users.Password, users.Email, users.PhoneNumber)
	return err
}

func UpdateUsers(db *sql.DB, users structs.Users) error {
	sql := "UPDATE users SET username=$1, password=$2, email=$3, phone_number=$4, updated_at=NOW() WHERE user_id=$5"

	_, err := db.Exec(sql, users.Username, users.Password, users.Email, users.PhoneNumber, users.UserId)

	if err != nil {
		return err
	}

	return nil
}

func DeleteUsers(db *sql.DB, users structs.Users) (err error) {
	sql := "DELETE FROM users WHERE user_id = $1"
	_, err = db.Exec(sql, users.UserId)
	return
}
