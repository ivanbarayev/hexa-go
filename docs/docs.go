// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "email": "ivanbarayev@hotmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/login": {
            "post": {
                "description": "Login process",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login",
                "parameters": [
                    {
                        "description": "` + "`" + `Body for user registration` + "`" + `",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.LoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.HandlerResponse"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "Registration process",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Register",
                "parameters": [
                    {
                        "description": "` + "`" + `Body for user registration` + "`" + `",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entities.RegisterReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entities.HandlerResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entities.HandlerResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "error": {
                    "type": "boolean"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "entities.LoginReq": {
            "type": "object",
            "required": [
                "src",
                "user_name",
                "user_pass",
                "user_type"
            ],
            "properties": {
                "src": {
                    "type": "integer",
                    "maximum": 5,
                    "minimum": 1
                },
                "user_name": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 10
                },
                "user_pass": {
                    "type": "string",
                    "maxLength": 16,
                    "minLength": 8
                },
                "user_type": {
                    "type": "integer",
                    "maximum": 5,
                    "minimum": 1
                },
                "verify_code": {
                    "type": "integer"
                }
            }
        },
        "entities.RegisterReq": {
            "type": "object",
            "required": [
                "src",
                "user_name",
                "user_phone",
                "user_title",
                "user_type"
            ],
            "properties": {
                "company_name": {
                    "type": "string"
                },
                "src": {
                    "type": "integer",
                    "maximum": 5,
                    "minimum": 1
                },
                "user_name": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 10
                },
                "user_pass": {
                    "type": "string"
                },
                "user_phone": {
                    "type": "string",
                    "maxLength": 20,
                    "minLength": 10
                },
                "user_title": {
                    "type": "string"
                },
                "user_type": {
                    "type": "integer",
                    "maximum": 5,
                    "minimum": 1
                },
                "verify_code": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/v1",
	Schemes:          []string{},
	Title:            "Auth Service",
	Description:      "Common Auth service broker with REST endpoints",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
