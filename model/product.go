package model

import "github.com/wintltr/vand-interview-crud-project/database"

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