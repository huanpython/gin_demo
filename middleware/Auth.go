/*
@Time : 2019/11/28 0028 下午 5:08
@Author : huanfuan
@File : Auth
@Software: GoLand
*/

package middleware

import "github.com/gin-gonic/gin"

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		println("已经授权")
		context.Next()
	}
}
