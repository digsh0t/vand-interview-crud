package model

import (
	"github.com/wintltr/vand-interview-crud-project/database"
)

type Store struct {
	StoreId int `json:"store_id"`
	Name string `json:"store_name"`
	Description string `json:"store_description"`
	UserId int `json:"store_userid"`
}

func InsertStoreToDB(store Store) error {
	db := database.ConnectDB()
	defer db.Close()

	stmt, err := db.Prepare(`INSERT INTO STORE (store_name, store_description, user_id) VALUES (?,?,?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(store.Name, store.Description, store.UserId)
	if err != nil {
		return err
	}
	return err
}
