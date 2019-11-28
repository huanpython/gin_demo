/*
@Time : 2019/11/28 0028 上午 11:41
@Author : huanfuan
@File : main
@Software: GoLand
*/

package main

import (
	"gin-demo/initRouter"
)

/*
https://youngxhui.top/2019/07/gin%E4%B8%80hello/
*/

func main() {
	router := initRouter.SetupRouter()
	_ = router.Run()

}
