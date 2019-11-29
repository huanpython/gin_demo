/*
@Time : 2019/11/29 0029 下午 4:52
@Author : huanfuan
@File : articleHandler
@Software: GoLand
*/

package article

import (
	"gin-demo/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func Insert(context *gin.Context) {
	article := model.Article{}
	var id = -1
	if e := context.ShouldBindJSON(&article); e == nil {
		id = article.Insert()
	}
	context.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func GetOne(context *gin.Context) {
	id := context.Param("id")
	i, e := strconv.Atoi(id)
	if e != nil {
		log.Panicln("id 不是 int 类型, id 转换失败", e.Error())
	}
	article := model.Article{
		Id: i,
	}

	art := article.FindById()
	context.JSON(http.StatusOK, gin.H{
		"article": art,
	})
}

func GetAll(context *gin.Context) {
	article := model.Article{}
	articles := article.FindAll()
	context.JSON(http.StatusOK, gin.H{
		"articles": articles,
	})
}

func DeleteOne(context *gin.Context) {
	id := context.Param("id")
	i, e := strconv.Atoi(id)
	if e != nil {
		log.Panicln("id 不是 int 类型, id 转换失败", e.Error())
	}
	article := model.Article{Id: i}
	article.DeleteOne()
}
