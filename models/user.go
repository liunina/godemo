/*
 * @Date: 2018-07-17 19:37:32
 * @LastEditors: liunian
 * @LastEditTime: 2020-09-14 20:36:52
 */
package models

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//user表结构体定义
type User struct {
	Id         int64  `json:"id" form:"id"`
	Username   string `json:"username" form:"username"`
	Password   string `json:"password" form:"password"`
	Status     int    `json:"status" form:"status"` // 0 正常状态， 1删除
	Createtime int64  `json:"createtime" form:"createtime"`
}
type JwtToken struct {
	Token string `json:"token"`
}

//添加
func InsertUser(user User) (id int64, err error) {
	//添加数据
	result := DB.Create(&user) //创建对象
	id = user.Id
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}
