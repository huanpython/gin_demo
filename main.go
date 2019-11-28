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

//https://youngxhui.top/2019/07/gin%E4%BA%8C%E8%B7%AF%E7%94%B1/

func main() {
	router := initRouter.SetupRouter()
	_ = router.Run()

}
