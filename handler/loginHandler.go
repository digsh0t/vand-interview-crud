package handler

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	simplejson "github.com/bitly/go-simplejson"
	"github.com/wintltr/vand-interview-crud-project/authentication"
	"github.com/wintltr/vand-interview-crud-project/database"
	"github.com/wintltr/vand-interview-crud-project/model"
	"github.com/wintltr/vand-interview-crud-project/util"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	var user model.User


	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		util.ERROR(w, http.StatusBadRequest, errors.New("Fail to read user data").Error())
		return
	}

	
	err = json.Unmarshal(body, &user)
	if err != nil {
		util.ERROR(w, http.StatusBadRequest, errors.New("Fail to parse json format").Error())
		return
	}

	err = model.InsertUserToDB(user)

	if err != nil {
		util.JSON(w, http.StatusBadRequest, nil)
	} else {
		util.JSON(w, http.StatusOK, user)
	}
}

func ListAllWebAppUser(w http.ResponseWriter, r *http.Request) {

	var waUserList []model.User
	waUserList, err := model.GetAllUserFromDB()

	// Return json
	if err != nil {
		util.ERROR(w, http.StatusBadRequest, errors.New("fail to get all list web app users").Error())
	} else {
		util.JSON(w, http.StatusOK, waUserList)
	}

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	var user model.User
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		util.ERROR(w, http.StatusBadRequest, "Username and Password is required")
	}

	json.Unmarshal(reqBody, &user)

	db := database.ConnectDB()
	defer db.Close()

	userId, err := model.CheckLogin(user.Username, user.Password)
	if err != nil {
		util.ERROR(w, http.StatusUnauthorized, "Wrong Username or Password")
		return
	} else {
		token, err := authentication.CreateToken(userId)
		if err != nil {
			util.ERROR(w, http.StatusUnauthorized, "Fail to create token while login")
			return
		}

		//Return Login Success Authorization Json
		returnJson := simplejson.New()
		returnJson.Set("Authorization", token)
		util.JSON(w, http.StatusOK, returnJson)
	}
}