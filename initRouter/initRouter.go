/*
@Time : 2019/11/28 0028 下午 1:59
@Author : huanfuan
@File : initRouter
@Software: GoLand
*/

package initRouter

import (
	"gin-demo/handler/article"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	articleRouter := router.Group("")
	{
		// 通过获取单篇文章
		articleRouter.GET("/article/:id", article.GetOne)
		// 获取所有文章
		articleRouter.GET("/articles", article.GetAll)
		// 添加一篇文章
		articleRouter.POST("/article", article.Insert)
		articleRouter.DELETE("/article/:id", article.DeleteOne)

	}

	return router

}
