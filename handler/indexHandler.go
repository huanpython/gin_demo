/*
@Time : 2019/11/28 0028 下午 3:11
@Author : huanfuan
@File : indexHandler
@Software: GoLand
*/

package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Index(context *gin.Context) {
	context.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "hello gin" + strings.ToLower(context.Request.Method) + "method",
	})
}
