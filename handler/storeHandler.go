package handler

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/wintltr/vand-interview-crud-project/authentication"
	"github.com/wintltr/vand-interview-crud-project/model"
	"github.com/wintltr/vand-interview-crud-project/util"
)

func AddStoreHandler(w http.ResponseWriter, r *http.Request) {

	var store model.Store

	tokenData, err := authentication.ExtractTokenMetadata(r)
	if err != nil {
		util.ERROR(w, http.StatusBadRequest, errors.New("Please login").Error())
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		util.ERROR(w, http.StatusBadRequest, errors.New("Fail to read user data").Error())
		return
	}
	
	err = json.Unmarshal(body, &store)
	if err != nil {
		util.ERROR(w, http.StatusBadRequest, errors.New("Fail to parse json format").Error())
		return
	}
	store.UserId = tokenData.Userid

	err = model.InsertStoreToDB(store)

	if err != nil {
		util.JSON(w, http.StatusBadRequest, errors.New(err.Error()))
	} else {
		util.JSON(w, http.StatusOK, store)
	}
}