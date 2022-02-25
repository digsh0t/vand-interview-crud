package model

import (
	"crypto/md5"
	"fmt"

	"github.com/wintltr/vand-interview-crud-project/database"
)

type User struct {
	UserId int `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
	Role string `json:"role"`
}

func InsertUserToDB(u User) error {
	db := database.ConnectDB()
	defer db.Close()

	stmt, err := db.Prepare(`INSERT INTO USER (user_username, user_password, user_email) VALUES (?,?,?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword := fmt.Sprintf("%x", md5.Sum([]byte(u.Password)))
	_, err = stmt.Exec(u.Username, hashedPassword, u.Email)
	if err != nil {
		return err
	}
	return err
}

func GetAllUserFromDB() ([]User, error) {
	db := database.ConnectDB()
	defer db.Close()

	var waUserList []User
	selDB, err := db.Query("SELECT user_id, user_username, user_role,user_email FROM USER")
	if err != nil {
		return waUserList, err
	}

	var user User
	for selDB.Next() {

		err = selDB.Scan(&user.UserId, &user.Username, &user.Role, &user.Email)
		if err != nil {
			return waUserList, err
		}
		waUserList = append(waUserList, user)
	}
	return waUserList, err
}

func CheckLogin(username string, password string) error {

	var userId int
	db := database.ConnectDB()
	defer db.Close()

	hashedPassword := fmt.Sprintf("%x", md5.Sum([]byte(password)))
	row := db.QueryRow("SELECT user_id FROM USER WHERE user_username = ? AND user_password = ?", username, hashedPassword)
	err := row.Scan(&userId)
	return err
}