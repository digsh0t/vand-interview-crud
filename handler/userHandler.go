package handler

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/wintltr/vand-interview-crud-project/model"
	"github.com/wintltr/vand-interview-crud-project/util"
)

func UserDetailHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	userId, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.ERROR(w, http.StatusBadRequest, "failed to get store detail")
		return
	}

	user, err := model.GetUserByIdFromDB(userId)

	// Return json
	if err != nil {
		util.ERROR(w, http.StatusBadRequest, err.Error())
	} else {
		util.JSON(w, http.StatusOK, user)
	}

}