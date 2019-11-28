/*
@Time : 2019/11/28 0028 下午 1:59
@Author : huanfuan
@File : initRouter
@Software: GoLand
*/

package initRouter

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello gin")
	})
	return router
}
