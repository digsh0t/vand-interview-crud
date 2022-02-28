package model

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"

	"github.com/wintltr/vand-interview-crud-project/database"
)

type Product struct {
	ProductId int `json:"product_id"`
	Name string `json:"product_name"`
	Price float64 `json:"product_price"`
	Variant string `json:"product_variant"`
	StoreId int `json:"store_id"`
}

func InsertProductToDB(product Product) error {
	db := database.ConnectDB()
	defer db.Close()

	stmt, err := db.Prepare(`INSERT INTO PRODUCT (product_name, product_price, product_variant,store_id) VALUES (?,?,?,?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Price, product.Variant, product.StoreId)
	if err != nil {
		return err
	}
	return err
}

func DeleteProductFromDB(id int) error {
	db := database.ConnectDB()
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM PRODUCT WHERE product_id = ?")
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

func GetAllProductFromDB() ([]Product, error) {
	db := database.ConnectDB()
	defer db.Close()

	var productList []Product
	selDB, err := db.Query("SELECT product_id, product_name, product_price, product_variant, store_id FROM PRODUCT")
	if err != nil {
		return productList, err
	}

	var product Product
	for selDB.Next() {
		err = selDB.Scan(&product.ProductId, &product.Name, &product.Price, &product.Variant, &product.StoreId)
		if err != nil {
			return productList, err
		}
		productList = append(productList, product)
	}

	return productList, err
}

func GetProductByIdFromDB(productId int) (Product, error) {
	db := database.ConnectDB()
	defer db.Close()

	var product Product
	row := db.QueryRow("SELECT product_id, product_name, product_price, product_variant, store_id FROM PRODUCT WHERE product_id = ?", productId)
	err := row.Scan(&product.ProductId, &product.Name, &product.Price, &product.Variant, &product.StoreId)
	if err != nil {
		return product, err
	}

	return product, err

}

func GetProductListByStoreFromDB(storeId int) ([]Product, error) {
	db := database.ConnectDB()
	defer db.Close()

	var productList []Product
	selDB, err := db.Query("SELECT product_id, product_name, product_price, product_variant, store_id FROM PRODUCT where store_id = ?",storeId)
	if err != nil {
		return productList, err
	}

	var product Product
	for selDB.Next() {
		err = selDB.Scan(&product.ProductId, &product.Name, &product.Price, &product.Variant, &product.StoreId)
		if err != nil {
			return productList, err
		}
		productList = append(productList, product)
	}

	return productList, err
}

func UpdateProductToDB(product Product) error {
	db := database.ConnectDB()
	defer db.Close()

	stmt, err := db.Prepare(`UPDATE PRODUCT SET product_name = ?, product_price = ?, product_variant = ? WHERE product_id = ?`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Price, product.Variant, product.ProductId)
	if err != nil {
		return err
	}
	return err
}

func CheckProductBelongUser(userId,productId int) (error) {
	db := database.ConnectDB()
	defer db.Close()

	product, err := GetProductByIdFromDB(productId)
	if err != nil {
		return err
	}

	err = CheckStoreBelongUser(userId, product.StoreId)
	return err
}

func GetTotalProductInDB() (int,error) {
	db := database.ConnectDB()
	defer db.Close()

	var total int
	row := db.QueryRow("SELECT COUNT(product_id) FROM PRODUCT")
	err := row.Scan(&total)
	if err != nil {
		return -1, err
	}

	return total, err
}

func GetProductByPage(page int, offset int) ([]Product, error) {
	var query string
	totalProduct, err := GetTotalProductInDB()
	if err != nil {
		return nil, err
	}
	if page * offset > totalProduct && (page-1) * offset > totalProduct {
		return nil, errors.New("The page exceed limit")
	} else {
		query = fmt.Sprintf(`SELECT product_id, product_name, product_price, product_variant, store_id FROM PRODUCT LIMIT %d OFFSET %d;`,offset, (page-1)*offset)
	}

	db := database.ConnectDB()
	defer db.Close()

	var productList []Product
	selDB, err := db.Query(query)
	if err != nil {
		return productList, err
	}

	var product Product
	for selDB.Next() {
		err = selDB.Scan(&product.ProductId, &product.Name, &product.Price, &product.Variant, &product.StoreId)
		if err != nil {
			return productList, err
		}
		productList = append(productList, product)
	}

	return productList, err
}

func SearchProductByPage(page int, offset int, searched string) ([]Product, error) {
	var query string
	totalProduct, err := GetTotalProductInDB()
	if err != nil {
		return nil, err
	}
	if page * offset > totalProduct && (page-1) * offset > totalProduct {
		return nil, errors.New("The page exceed limit")
	} else {
		query = fmt.Sprintf(`SELECT product_id, product_name, product_price, product_variant, store_id FROM PRODUCT WHERE product_name LIKE '%%%s%%' LIMIT %d OFFSET %d;`,searched, offset, (page-1)*offset)
	}

	db := database.ConnectDB()
	defer db.Close()

	var productList []Product
	selDB, err := db.Query(query)
	if err != nil {
		return productList, err
	}

	var product Product
	for selDB.Next() {
		err = selDB.Scan(&product.ProductId, &product.Name, &product.Price, &product.Variant, &product.StoreId)
		if err != nil {
			return productList, err
		}
		productList = append(productList, product)
	}

	return productList, err
}

func ExportAllProductToCSV(location string) error {
	productList, err := GetAllProductFromDB()
	if err != nil {
		return err
	}
	csvFile, err := os.Create(location)

	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)

	for _, product := range productList {
		var row []string
		row = append(row,fmt.Sprint(product.ProductId))
		row = append(row, product.Name)
		row = append(row, fmt.Sprintf("%.3f", product.Price))
		row = append(row, product.Variant)
		row = append(row,fmt.Sprint(product.StoreId))
		writer.Write(row)
	}

	// remember to flush!
	writer.Flush()
	return err
}