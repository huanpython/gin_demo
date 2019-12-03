/*
@Time : 2019/12/3 0003 上午 11:30
@Author : huanfuan
@File : role
@Software: GoLand
*/
package model

import (
	"gin-demo/initDB"
	"github.com/jinzhu/gorm"
)

type Role struct {
	gorm.Model
	Name string `json:"name"`
}

func init() {
	if !initDB.Db.HasTable(Role{}) {
		initDB.Db.CreateTable(Role{})
	}
}
