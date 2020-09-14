/*
 * @Date: 2018-07-17 19:37:32
 * @LastEditors: liunian
 * @LastEditTime: 2020-09-14 17:48:26
 */
package main

import (
	"net/http"

	"github.com/liunina/godemo/routes"
)

func main() {
	r := routes.NewRouter()

	http.ListenAndServe(":8080", r)
}
