/*
@Time : 2019/11/28 0028 下午 1:59
@Author : huanfuan
@File : initRouter
@Software: GoLand
*/

package initRouter

import (
	"gin-demo/handler/user"
	"gin-demo/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", middleware.Auth(), func(context *gin.Context) {
		context.JSON(http.StatusOK, time.Now().Unix())
	})
	router.GET("/login", user.CreateJwt)
	router.POST("/register", user.Register)

	return router

}
