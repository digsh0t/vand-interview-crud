package handler

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/wintltr/vand-interview-crud-project/authentication"
	"github.com/wintltr/vand-interview-crud-project/model"
	"github.com/wintltr/vand-interview-crud-project/util"
)

func AddProductHandler(w http.ResponseWriter, r *http.Request) {

	var product model.Product

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		util.ERROR(w, http.StatusBadRequest, errors.New("Fail to read user data").Error())
		return
	}
	
	err = json.Unmarshal(body, &product)
	if err != nil {
		util.ERROR(w, http.StatusBadRequest, errors.New("Fail to parse json format").Error())
		return
	}

	err = model.InsertProductToDB(product)

	if err != nil {
		util.JSON(w, http.StatusBadRequest, errors.New(err.Error()))
	} else {
		util.JSON(w, http.StatusOK, product)
	}
}

func RemoveProductHandler(w http.ResponseWriter, r *http.Request) {

	tokenData, err := authentication.ExtractTokenMetadata(r)
	if err != nil {
		util.ERROR(w, http.StatusBadRequest, errors.New("Please login").Error())
		return
	}

	// Retrieve Id
	vars := mux.Vars(r)
	productId, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.ERROR(w, http.StatusUnauthorized, errors.New("Failed to retrieve Id").Error())
		return
	}

	//Check authorization
	err = model.CheckProductBelongUser(tokenData.Userid, productId)
	if err != nil {
		util.ERROR(w, http.StatusBadRequest, errors.New("You are not authorized to modify this store").Error())
		return
	}

	err = model.DeleteProductFromDB(productId)
	if err != nil {
		util.ERROR(w, http.StatusOK, errors.New("Failed to remove product from database").Error())
	} else {
		util.JSON(w, http.StatusOK, nil)
	}
}

func ListProductByPageHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	page, err := strconv.Atoi(vars["page"])
	if err != nil {
		util.ERROR(w, http.StatusBadRequest, "failed to get page value")
		return
	}

	productList, err := model.GetProductByPage(page, 10)

	// Return json
	if err != nil {
		util.ERROR(w, http.StatusBadRequest, errors.New("fail to get products").Error())
	} else {
		util.JSON(w, http.StatusOK, productList)
	}
}

func SearchProductByPageHandler(w http.ResponseWriter, r *http.Request) {

	type received struct {
		SearchString string `json:"search_string"`
		Page int `json:"page"`
	}

	var receivedValue received

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		util.ERROR(w, http.StatusBadRequest, errors.New("Fail to read user data").Error())
		return
	}
	
	err = json.Unmarshal(body, &receivedValue)
	if err != nil {
		util.ERROR(w, http.StatusBadRequest, errors.New("Fail to parse json format").Error())
		return
	}

	if receivedValue.Page == 0 {
		receivedValue.Page = 1
	}

	productList, err := model.SearchProductByPage(receivedValue.Page, 10, receivedValue.SearchString)

	// Return json
	if err != nil {
		util.ERROR(w, http.StatusBadRequest, errors.New("fail to get products").Error())
	} else {
		util.JSON(w, http.StatusOK, productList)
	}
}

func ProductDetailHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	productId, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.ERROR(w, http.StatusBadRequest, "failed to get product detail")
		return
	}

	product, err := model.GetProductByIdFromDB(productId)

	// Return json
	if err != nil {
		util.ERROR(w, http.StatusBadRequest, err.Error())
	} else {
		util.JSON(w, http.StatusOK, product)
	}

}

func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {

	var product model.Product

	tokenData, err := authentication.ExtractTokenMetadata(r)
	if err != nil {
		util.ERROR(w, http.StatusBadRequest, errors.New("Please login").Error())
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		util.ERROR(w, http.StatusBadRequest, errors.New("Fail to update product").Error())
		return
	}
	
	err = json.Unmarshal(body, &product)
	if err != nil {
		util.ERROR(w, http.StatusBadRequest, errors.New("Fail to parse json format").Error())
		return
	}

	//Check authorization
	err = model.CheckProductBelongUser(tokenData.Userid, product.ProductId)
	if err != nil {
		util.ERROR(w, http.StatusBadRequest, errors.New("You are not authorized to modify this store").Error())
		return
	}

	err = model.UpdateProductToDB(product)

	if err != nil {
		util.JSON(w, http.StatusBadRequest, errors.New(err.Error()))
	} else {
		util.JSON(w, http.StatusOK, product)
	}
}