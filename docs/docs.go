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
        "/api/healthz": {
            "get": {
                "description": "Get Healthz",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Healthz"
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/v1/api_key": {
            "get": {
                "description": "GetApiKey",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "GetApiKey"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Api Key",
                        "name": "api_key",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "put": {
                "description": "UpdateApiKey",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "UpdateApiKey"
                ],
                "parameters": [
                    {
                        "description": "UploadApiKeyRequest Model",
                        "name": "uploadApiKeyRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UploadApiKeyRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "post": {
                "description": "AddApiKey",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "AddApiKey"
                ],
                "parameters": [
                    {
                        "description": "UploadApiKeyRequest Model",
                        "name": "uploadApiKeyRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UploadApiKeyRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "delete": {
                "description": "DeleteApiKey",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "DeleteApiKey"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Api Key",
                        "name": "api_key",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/v1/api_key/list": {
            "get": {
                "description": "GetApiKeyList",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "GetApiKeyList"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/v1/paymaster_strategy": {
            "get": {
                "description": "GetStrategy",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "GetStrategy"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "PaymasterStrategy Code",
                        "name": "strategy_code",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "put": {
                "description": "UpdateStrategy",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "UpdateStrategy"
                ],
                "parameters": [
                    {
                        "description": "UploadStrategyRequest Model",
                        "name": "uploadStrategyRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UploadStrategyRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "post": {
                "description": "AddStrategy",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "AddStrategy"
                ],
                "parameters": [
                    {
                        "description": "UploadStrategyRequest Model",
                        "name": "uploadStrategyRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UploadStrategyRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "delete": {
                "description": "DeleteStrategy",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "DeleteStrategy"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "PaymasterStrategy Code",
                        "name": "strategy_code",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/v1/paymaster_strategy/list": {
            "get": {
                "description": "GetStrategyList",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "GetStrategyList"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "model.UploadApiKeyRequest": {
            "type": "object",
            "properties": {
                "api_key": {
                    "type": "string"
                },
                "api_key_name": {
                    "type": "string"
                },
                "extra_info": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "model.UploadStrategyRequest": {
            "type": "object",
            "properties": {
                "extra": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "project_code": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "strategy_code": {
                    "type": "string"
                },
                "strategy_execute_restriction": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "strategy_name": {
                    "type": "string"
                },
                "user_id": {
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
