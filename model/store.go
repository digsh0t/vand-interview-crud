package model

import (
	"errors"
	"fmt"

	"github.com/wintltr/vand-interview-crud-project/database"
)

type Store struct {
	StoreId int `json:"store_id"`
	Name string `json:"store_name"`
	Description string `json:"store_description"`
	ProductList []Product `json:"product_list"`
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

func DeleteStoreFromDB(id int) error {
	db := database.ConnectDB()
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM STORE WHERE store_id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if rows == 0 {
		return err
	}
	return err
}

func GetAllStoreFromDB() ([]Store, error) {
	db := database.ConnectDB()
	defer db.Close()

	var storeList []Store
	selDB, err := db.Query("SELECT store_id, store_name, store_description, user_id FROM STORE")
	if err != nil {
		return storeList, err
	}

	var store Store
	for selDB.Next() {
		err = selDB.Scan(&store.StoreId, &store.Name, &store.Description, &store.UserId)
		if err != nil {
			return storeList, err
		}
		storeList = append(storeList, store)
	}

	return storeList, err
}

// func GetAllStoreFromDBByPage(page int) ([]Store, error) {
// 	storeList, err := GetAllStoreFromDB()
// 	if err != nil {
// 		return nil, err
// 	}
	
// }

func GetStoreByIdFromDB(storeId int) (Store, error) {
	db := database.ConnectDB()
	defer db.Close()

	var store Store
	row := db.QueryRow("SELECT store_id, store_name, store_description, user_id FROM STORE WHERE store_id = ?", storeId)
	err := row.Scan(&store.StoreId, &store.Name, &store.Description, &store.UserId)
	if err != nil {
		return store, err
	}

	store.ProductList, err = GetProductListByStoreFromDB(store.StoreId)
	return store, err
}

func GetStoreListByUserFromDB(userId int) ([]Store, error) {
	db := database.ConnectDB()
	defer db.Close()

	var storeList []Store
	selDB, err := db.Query("SELECT store_id, store_name, store_description, user_id FROM STORE where user_id = ?",userId)
	if err != nil {
		return storeList, err
	}

	var store Store
	for selDB.Next() {
		err = selDB.Scan(&store.StoreId, &store.Name, &store.Description, &store.UserId)
		if err != nil {
			return storeList, err
		}
		store.ProductList, err = GetProductListByStoreFromDB(store.StoreId)
		if err != nil {
			return storeList, err
		}
		storeList = append(storeList, store)
	}

	return storeList, err
}

func UpdateStoreToDB(store Store) error {
	db := database.ConnectDB()
	defer db.Close()

	stmt, err := db.Prepare(`UPDATE STORE SET store_name = ?, store_description = ? WHERE store_id = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(store.Name, store.Description, store.StoreId)
	if err != nil {
		return err
	}
	return err
}

func CheckStoreBelongUser(userId,storeId int) (error) {
	db := database.ConnectDB()
	defer db.Close()

	row, err := db.Query("SELECT store_id FROM STORE where user_id = ? AND store_id = ?",userId, storeId)
	if err != nil || !row.Next() {
		return errors.New("No row returned")
	}
	return err
}

func GetTotalStoreInDB() (int,error) {
	db := database.ConnectDB()
	defer db.Close()

	var total int
	row := db.QueryRow("SELECT COUNT(store_id) FROM STORE")
	err := row.Scan(&total)
	if err != nil {
		return -1, err
	}

	return total, err
}

func GetStoreByPage(page int, offset int) ([]Store, error) {
	var query string
	totalStore, err := GetTotalStoreInDB()
	if err != nil {
		return nil, err
	}
	if page * offset > totalStore && (page-1) * offset > totalStore {
		return nil, errors.New("The page exceed limit")
	} else {
		query = fmt.Sprintf(`SELECT store_id, store_name, store_description, user_id FROM STORE LIMIT %d OFFSET %d;`,offset, (page-1)*offset)
	}

	db := database.ConnectDB()
	defer db.Close()

	var storeList []Store
	selDB, err := db.Query(query)
	if err != nil {
		return storeList, err
	}

	var store Store
	for selDB.Next() {
		err = selDB.Scan(&store.StoreId, &store.Name, &store.Description, &store.UserId)
		if err != nil {
			return storeList, err
		}
		store.ProductList, err = GetProductListByStoreFromDB(store.StoreId)
		if err != nil {
			return storeList, err
		}
		storeList = append(storeList, store)
	}

	return storeList, err
}

func SearchStoreByPage(page int, offset int, searched string) ([]Store, error) {
	var query string
	totalStore, err := GetTotalStoreInDB()
	if err != nil {
		return nil, err
	}
	if page * offset > totalStore && (page-1) * offset > totalStore {
		return nil, errors.New("The page exceed limit")
	} else {
		query = fmt.Sprintf(`SELECT store_id, store_name, store_description, user_id FROM STORE WHERE store_name LIKE '%%%s%%' LIMIT %d OFFSET %d;`,searched, offset, (page-1)*offset)
	}

	db := database.ConnectDB()
	defer db.Close()

	var storeList []Store
	selDB, err := db.Query(query)
	if err != nil {
		return storeList, err
	}

	var store Store
	for selDB.Next() {
		err = selDB.Scan(&store.StoreId, &store.Name, &store.Description, &store.UserId)
		if err != nil {
			return storeList, err
		}
		store.ProductList, err = GetProductListByStoreFromDB(store.StoreId)
		if err != nil {
			return storeList, err
		}
		storeList = append(storeList, store)
	}

	return storeList, err
}