/*
 * @Date: 2018-07-17 19:37:32
 * @LastEditors: liunian
 * @LastEditTime: 2020-09-14 23:05:11
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
func (user User) Check() (id int64, err error) {
	result := DB.Where("username = ?", user.Username).First(&user)
	id = user.Id
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

//添加
func (user User) Insert() (id int64, err error) {
	existId, err := user.Check()
	if err == nil || existId > 0 {
		return
	}

	//添加数据
	result := DB.Create(&user) //创建对象
	id = user.Id
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

func (user User) FindId() (userInfo User, err error) {
	result := DB.Where("id = ?", user.Id).First(&user)

	if result.Error != nil {
		err = result.Error
		return
	}
	userInfo = user
	return
}
