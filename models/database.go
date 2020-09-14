/*
 * @Date: 2020-09-14 17:40:07
 * @LastEditors: liunian
 * @LastEditTime: 2020-09-14 20:55:49
 */

package models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" //加载mysql
	"github.com/jinzhu/gorm"
)

//数据库连接信息
const (
	USERNAME = "liunian"
	PASSWORD = "hongye"
	NETWORK  = "tcp"
	SERVER   = "10.1.1.2"
	PORT     = 3306
	DATABASE = "test"
)

var DB *gorm.DB

func init() {

	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	var err error
	DB, err = gorm.Open("mysql", conn)
	if err != nil {
		fmt.Println("failed to connect database:", err)
		return
	} else {
		fmt.Println("connect database success")
		DB.SingularTable(true)
		DB.AutoMigrate(&User{}) //自动建表
		fmt.Println("create table success")
	}

	// defer DB.Close()
}
