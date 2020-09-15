/*
 * @Date: 2018-07-17 19:37:32
 * @LastEditors: liunian
 * @LastEditTime: 2020-09-15 10:19:40
 */
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/liunina/godemo/routes"
)

func main() {
	logFile, err := os.OpenFile("./log/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	r := routes.NewRouter()

	http.ListenAndServe(":8081", r)
}
