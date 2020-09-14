/*
 * @Date: 2018-07-17 19:37:32
 * @LastEditors: liunian
 * @LastEditTime: 2020-09-14 21:48:35
 */
package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ResponseWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

type BusError struct {
	Code    int
	Message string
}

func (e *BusError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}
