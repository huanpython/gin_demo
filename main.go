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

//https://youngxhui.top/2019/07/gin%E4%B9%9D%E7%94%9F%E6%88%90restful%E6%8E%A5%E5%8F%A3/

func main() {
	router := initRouter.SetupRouter()
	_ = router.Run()

}
