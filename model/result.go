/*
@Time : 2019/12/3 0003 上午 10:45
@Author : huanfuan
@File : Result
@Software: GoLand
*/
package model

type Result struct {
	Code    int         `json:"code" example:"000"`
	Message string      `json:"message" example:"请求信息"`
	Data    interface{} `json:"data" `
}
