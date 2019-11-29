/*
@Time : 2019/11/28 0028 下午 4:06
@Author : huanfuan
@File : initDB
@Software: GoLand
*/

package initDB

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/ginhello")
	if err != nil {
		log.Panicln("数据库连接失败:", err.Error())
	}
	Db.SetMaxOpenConns(20)
	Db.SetMaxIdleConns(20)
}
