/*
 * @Date: 2018-07-17 19:37:32
 * @LastEditors: liunian
 * @LastEditTime: 2020-09-14 19:00:33
 */
package main

import (
	"net/http"

	"github.com/liunina/godemo/routes"
)

func main() {
	r := routes.NewRouter()

	http.ListenAndServe(":8081", r)
}
