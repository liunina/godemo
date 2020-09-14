/*
 * @Date: 2018-07-17 19:37:32
 * @LastEditors: liunian
 * @LastEditTime: 2020-09-14 20:48:50
 */
package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/liunina/godemo/auth"
	"github.com/liunina/godemo/helper"
	"github.com/liunina/godemo/models"
)

const (
	db         = "Movies"
	collection = "UserModel"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var jsonStr = json.NewDecoder(r.Body)
	fmt.Println(jsonStr)
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil || user.Username == "" || user.Password == "" {
		helper.ResponseWithJson(w, http.StatusBadRequest,
			helper.Response{Code: http.StatusBadRequest, Msg: "bad params"})
		return
	}
	id, err := user.Insert()
	if err != nil {
		helper.ResponseWithJson(w, http.StatusInternalServerError,
			helper.Response{Code: http.StatusInternalServerError, Msg: "internal error"})
	} else {
		helper.ResponseWithJson(w, http.StatusBadRequest,
			helper.Response{Code: http.StatusOK, Data: id})
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helper.ResponseWithJson(w, http.StatusBadRequest,
			helper.Response{Code: http.StatusBadRequest, Msg: "bad params"})
	}
	// exist := models.IsExist(db, collection, bson.M{"username": user.UserName})

	var exist = false

	if exist {
		token, _ := auth.GenerateToken(&user)
		helper.ResponseWithJson(w, http.StatusOK,
			helper.Response{Code: http.StatusOK, Data: models.JwtToken{Token: token}})
	} else {
		helper.ResponseWithJson(w, http.StatusNotFound,
			helper.Response{Code: http.StatusNotFound, Msg: "the user not exist"})
	}
}
