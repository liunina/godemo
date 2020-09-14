/*
 * @Date: 2018-07-17 19:37:32
 * @LastEditors: liunian
 * @LastEditTime: 2020-09-14 21:46:11
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

var Users []User

// 检测查询
func (user User) Check() (id int64, err BusError) {
	result := DB.Where("username = ?", user.Username).First(&user)
	id = user.Id
	if result.Error != nil {
		err = &BusError(-1, "发送错误")
		return
	}
	return
}

//添加
func (user User) Insert() (id int64, err BusError) {
	id, err0 := user.Check()
	if err0 != nil {
		err = &BusError(-1, "发送错误")
		return
	}

	if id > 0 {
		err = &BusError(-1, "发送错误")
		return
	}

	//添加数据
	result := DB.Create(&user) //创建对象
	id = user.Id
	if result.Error != nil {
		err = &BusError(-1, "发送错误")
		return
	}
	return
}
