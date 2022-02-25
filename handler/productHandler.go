package handler

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

	// Retrieve Id
	vars := mux.Vars(r)
	storeId, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.ERROR(w, http.StatusUnauthorized, errors.New("Failed to retrieve Id").Error())
		return
	}

	err = model.DeleteProductFromDB(storeId)
	if err != nil {
		util.ERROR(w, http.StatusOK, errors.New("Failed to remove product from database").Error())
	} else {
		util.JSON(w, http.StatusOK, nil)
	}
}

func ListAllProductHandler(w http.ResponseWriter, r *http.Request) {

	productList, err := model.GetAllProductFromDB()

	// Return json
	if err != nil {
		util.ERROR(w, http.StatusBadRequest, errors.New("fail to get all list web app users").Error())
	} else {
		util.JSON(w, http.StatusOK, productList)
	}
}