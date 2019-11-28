/*
@Time : 2019/11/28 0028 下午 2:02
@Author : huanfuan
@File : index_test
@Software: GoLand
*/

package test

import (
	"gin-demo/initRouter"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var router *gin.Engine

func init() {
	gin.SetMode(gin.TestMode)
	router = initRouter.SetupRouter()
}

func TestIndexHtml(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "hello gin get method", "返回的HTML页面中应该包含 hello gin get method")
}
