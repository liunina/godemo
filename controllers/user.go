/*
 * @Date: 2018-07-17 19:37:32
 * @LastEditors: liunian
 * @LastEditTime: 2020-09-14 22:55:17
 */
package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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

	existId, err := user.Check()
	if existId > 0 {
		helper.ResponseWithJson(w, http.StatusBadRequest,
			helper.Response{Code: http.StatusBadRequest, Msg: "user is exist!!"})
		return
	}

	id, err := user.Insert()
	fmt.Println(id)
	if err != nil {
		helper.ResponseWithJson(w, http.StatusInternalServerError,
			helper.Response{Code: http.StatusInternalServerError, Msg: "internal error"})
	} else {
		helper.ResponseWithJson(w, http.StatusBadRequest,
			helper.Response{Code: http.StatusOK, Data: &user})
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		helper.ResponseWithJson(w, http.StatusBadRequest,
			helper.Response{Code: http.StatusBadRequest, Msg: "bad params"})
	}

	id, err := user.Check()

	var exist = false
	if id > 0 {
		exist = true
	}

	if exist {
		token, _ := auth.GenerateToken(&user)
		helper.ResponseWithJson(w, http.StatusOK,
			helper.Response{Code: http.StatusOK, Data: models.JwtToken{Token: token}})
	} else {
		helper.ResponseWithJson(w, http.StatusNotFound,
			helper.Response{Code: http.StatusNotFound, Msg: "the user not exist"})
	}
}

func UserInfo(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	queryId := query.Get("id")
	var id int64 = 0

	id, err := strconv.ParseInt(queryId, 10, 64)
	if err == nil {
		var user models.User
		user.Id = id

		userInfo, err := user.FindId()
		if err != nil {
			helper.ResponseWithJson(w, http.StatusInternalServerError,
				helper.Response{Code: http.StatusInternalServerError, Msg: "internal error"})
		} else {
			helper.ResponseWithJson(w, http.StatusBadRequest,
				helper.Response{Code: http.StatusOK, Data: &userInfo})
		}
	} else {
		helper.ResponseWithJson(w, http.StatusBadRequest,
			helper.Response{Code: http.StatusBadRequest, Msg: "bad params"})
		return
	}

}
