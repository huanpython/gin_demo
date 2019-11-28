/*
@Time : 2019/11/28 0028 下午 1:59
@Author : huanfuan
@File : initRouter
@Software: GoLand
*/

package initRouter

import (
	"gin-demo/handler"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	if mode := gin.Mode(); mode == gin.TestMode {
		router.LoadHTMLGlob("./../templates/*")
	} else {
		router.LoadHTMLGlob("templates/*")
	}

	router.StaticFile("/favicon.ico", "./favicon.ico")
	router.Static("/statics", "./statics")
	index := router.Group("/")
	{
		index.Any("", handler.Index)
	}

	// 添加 user
	userRouter := router.Group("/user")
	{
		userRouter.GET("/:name", handler.UserSave)
		userRouter.GET("", handler.UserSaveByQuery)
		userRouter.POST("/register", handler.UserRegister)
	}
	return router

}
