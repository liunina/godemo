/*
 * @Date: 2018-07-17 19:37:32
 * @LastEditors: liunian
 * @LastEditTime: 2020-09-14 17:45:33
 */
package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	Id       int    `gorm:"size:11;primary_key;AUTO_INCREMENT;not null" json:"id"`
	UserName string `gorm:"size:255;DEFAULT NULL" json:"username"`
	Password string `gorm:"size:255;DEFAULT NULL" json:"password"`
	//gorm后添加约束，json后为对应mysql里的字段
}
type JwtToken struct {
	Token string `json:"token"`
}
