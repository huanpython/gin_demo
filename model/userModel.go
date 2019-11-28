/*
@Time : 2019/11/28 0028 下午 3:35
@Author : huanfuan
@File : userModel
@Software: GoLand
*/

package model

type UserModel struct {
	Email         string `form:"email" binding:"eamil" `
	Password      string `form:"password"`
	PasswordAgain string `form:"password-again" binding:"eqfield=Password"`
}
