/*
@Time : 2019/12/3 0003 下午 2:59
@Author : huanfuan
@File : docs
@Software: GoLand
*/

package docs

import (
	"bytes"
	"encoding/json"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "youngxhu",
            "url": "https://youngxhui.top",
            "email": "youngxhui@g mail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/article": {
            "post": {
                "consumes": [
                    "application/x-json-stream"
                ],
                "tags": [
                    "文章"
                ],
                "summary": "提交新的文章内容",
                "operationId": "1",
                "parameters": [
                    {
                        "description": "文章",
                        "name": "article",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/model.Article"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Result"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/model.Result"
                        }
                    }
                }
            }
        },
        "/article/{id}": {
            "get": {
                "consumes": [
                    "application/x-json-stream"
                ],
                "summary": "通过文章 id 获取单个文章内容",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Result"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/x-json-stream"
                ],
                "summary": "通过id删除指定文章类型",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Result"
                        }
                    }
                }
            }
        },
        "/articles": {
            "get": {
                "consumes": [
                    "application/x-json-stream"
                ],
                "summary": "获取所有的文章",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Result"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Article": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string",
                    "example": "hello js"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "type": {
                    "type": "string",
                    "example": "js"
                }
            }
        },
        "model.Result": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 0
                },
                "data": {
                    "type": "object"
                },
                "message": {
                    "type": "string",
                    "example": "请求信息"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "localhost:8080",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Gin swagger",
	Description: "Gin swagger 示例项目",
}

type s struct{}

func (s *s) ReadDoc() string {
	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, SwaggerInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
