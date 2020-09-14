/*
 * @Date: 2020-09-14 17:40:07
 * @LastEditors: liunian
 * @LastEditTime: 2020-09-14 17:44:48
 */

package models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" //加载mysql
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

const (
	host = "10.1.1.2:3306"
	user = "liunian"
	pass = "hongye"
)

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:root@tcp(10.1.1.2:3306)/test?charset=utf8&parseTime=True&loc=Local&timeout=10ms")

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	if db.Error != nil {
		fmt.Printf("database error %v", db.Error)
	}
}
