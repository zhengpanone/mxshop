// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/user/list": {
            "get": {
                "description": "获取用户列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "用户列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token令牌",
                        "name": "x-token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 1,
                        "description": "页码",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "default": 10,
                        "description": "页面大小",
                        "name": "size",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "Errors": {
                                            "type": "object"
                                        },
                                        "data": {
                                            "type": "object"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/user/pwd_login": {
            "post": {
                "description": "用户账号密码登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "请求参数",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/forms.PasswordLoginForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "object"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/v1/user/register": {
            "post": {
                "description": "用户注册",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "用户注册",
                "parameters": [
                    {
                        "description": "请求参数",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/forms.RegisterForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/utils.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "object"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "forms.PasswordLoginForm": {
            "type": "object",
            "required": [
                "captcha",
                "captcha_id",
                "mobile",
                "password"
            ],
            "properties": {
                "captcha": {
                    "type": "string",
                    "maxLength": 5,
                    "minLength": 5
                },
                "captcha_id": {
                    "type": "string"
                },
                "mobile": {
                    "description": "手机号码 自定义validator",
                    "type": "string"
                },
                "password": {
                    "description": "密码",
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 3
                }
            }
        },
        "forms.RegisterForm": {
            "type": "object",
            "required": [
                "code",
                "mobile",
                "password"
            ],
            "properties": {
                "code": {
                    "description": "短信验证码",
                    "type": "string",
                    "maxLength": 5,
                    "minLength": 5
                },
                "mobile": {
                    "description": "手机号码 自定义validator",
                    "type": "string"
                },
                "password": {
                    "description": "密码",
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 3
                }
            }
        },
        "utils.ErrorItem": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "key": {
                    "type": "string"
                }
            }
        },
        "utils.Meta": {
            "type": "object",
            "properties": {
                "request_id": {
                    "type": "string"
                }
            }
        },
        "utils.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "description": "业务状态码",
                    "type": "integer"
                },
                "data": {
                    "description": "响应数据"
                },
                "errors": {
                    "description": "Errors 错误提示，如 xx字段不能为空等",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/utils.ErrorItem"
                    }
                },
                "meta": {
                    "description": "Meta 源数据,存储如请求ID,分页等信息",
                    "allOf": [
                        {
                            "$ref": "#/definitions/utils.Meta"
                        }
                    ]
                },
                "msg": {
                    "description": "提示信息",
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
