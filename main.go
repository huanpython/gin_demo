/*
@Time : 2019/11/28 0028 上午 11:41
@Author : huanfuan
@File : main
@Software: GoLand
*/

package main

import (
	_ "gin-demo/docs"
	"gin-demo/initRouter"
)

// @title Gin swagger
// @version 1.0
// @description Gin swagger 示例项目

// @contact.name youngxhu
// @contact.url https://youngxhui.top
// @contact.email youngxhui@g mail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080

func main() {
	router := initRouter.SetupRouter()
	_ = router.Run()

}
