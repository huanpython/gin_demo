/*
@Time : 2019/12/3 0003 上午 10:49
@Author : huanfuan
@File : comment
@Software: GoLand
*/
package model

import (
	"gin-demo/initDB"
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	Content string
}

func init() {
	table := initDB.Db.HasTable(Comment{})
	if !table {
		initDB.Db.CreateTable(Comment{})
	}
}
