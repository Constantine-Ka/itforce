// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Константин",
            "url": "https://t.me/London68"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/history": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Показывает историю по одному или всем пользователям",
                "operationId": "pistory-handler",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Идентификатор пользователя",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/repositories.ActionHistory"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorShort"
                        }
                    }
                }
            }
        },
        "/purchase": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Изменяет баланс. Использует Мьютексы, для блокировки параллельного изменения данных",
                "operationId": "purchase-handler",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Идентификатор пользователя",
                        "name": "user",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Сумма списания",
                        "name": "amount",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorMsg"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/api.ErrorMsg"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "api.ErrorMsg": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "string"
                },
                "balance": {
                    "type": "number"
                },
                "historyId": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "mode": {
                    "type": "string"
                },
                "userId": {
                    "type": "string"
                }
            }
        },
        "api.ErrorShort": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "userId": {
                    "type": "string"
                }
            }
        },
        "repositories.ActionHistory": {
            "type": "object",
            "properties": {
                "actionEng": {
                    "type": "string"
                },
                "actionRus": {
                    "type": "string"
                },
                "amount": {
                    "type": "number"
                },
                "currentbalance": {
                    "type": "number"
                },
                "id": {
                    "type": "integer"
                },
                "ts": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "itForce",
	Description:      "Тестовое задание.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
