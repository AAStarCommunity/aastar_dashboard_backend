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
            "name": "AAStar BackEndDashBoard",
            "url": "https://aastar.xyz"
        },
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
                "security": [
                    {
                        "JWT": []
                    }
                ],
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
                "security": [
                    {
                        "JWT": []
                    }
                ],
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
            "delete": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
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
        "/api/v1/api_key/apply": {
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "ApplyApiKey",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "ApplyApiKey"
                ],
                "parameters": [
                    {
                        "description": "UploadApiKeyRequest Model",
                        "name": "applyApiKeyRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ApplyApiKeyRequest"
                        }
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
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "GetApiKeyList",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "GetApiKeyList"
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/v1/data/api_keys_data_overview": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "DataViewApiKeysOverView",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "DataViewApiKeysOverView"
                ],
                "summary": "DataViewApiKeysOverView",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/v1/data/balance": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "DataViewGetBalance",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "DataViewGetBalance"
                ],
                "summary": "DataViewGetBalance",
                "parameters": [
                    {
                        "type": "boolean",
                        "description": "is_test_net",
                        "name": "is_test_net",
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
        "/api/v1/data/paymaster_pay_type_metrics": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "DataViewPaymasterPayTypeMetrics",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "DataViewPaymasterPayTypeMetrics"
                ],
                "summary": "DataViewPaymasterPayTypeMetrics",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/v1/data/paymaster_requests": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "DataViewApiKeyPaymasterRecallDetailList",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "DataViewApiKeyPaymasterRecallDetailList"
                ],
                "summary": "DataViewApiKeyPaymasterRecallDetailList",
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
            }
        },
        "/api/v1/data/request_health_list": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "DataViewRequestHealth",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "DataViewRequestHealth"
                ],
                "summary": "DataViewRequestHealth",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Api Key",
                        "name": "api_key",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/v1/data/request_health_one": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "DataViewRequestHealthOneByApiKey",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "DataViewRequestHealthOneByApiKey"
                ],
                "summary": "DataViewRequestHealthOneByApiKey",
                "parameters": [
                    {
                        "type": "string",
                        "description": "API Key",
                        "name": "api_key",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Time Type",
                        "name": "time_type",
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
        "/api/v1/data/sponsor_transaction_list": {
            "get": {
                "description": "DataViewGetSponsorTransactionList",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "DataViewGetSponsorTransactionList"
                ],
                "summary": "DataViewGetSponsorTransactionList",
                "parameters": [
                    {
                        "type": "boolean",
                        "description": "is_test_net",
                        "name": "is_test_net",
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
        "/api/v1/data/sponsored_metrics": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "DataViewSponsoredMetrics",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "DataViewSponsoredMetrics"
                ],
                "summary": "DataViewSponsoredMetrics",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/v1/paymaster_strategy": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
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
                "security": [
                    {
                        "JWT": []
                    }
                ],
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
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
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
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
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
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "GetStrategyList",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "GetStrategyList"
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/v1/paymaster_strategy/switch_status": {
            "put": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "SwitchStrategyStatus",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "SwitchStrategyStatus"
                ],
                "parameters": [
                    {
                        "description": "ChangeStrategyStatusRequest Model",
                        "name": "ChangeStrategyStatusRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.ChangeStrategyStatusRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/api/v1/user": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "GetUserInfo",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "GetUserInfo"
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/oauth/github": {
            "get": {
                "description": "Github OAuth Login",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "OAuth"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "Github OAuth Code",
                        "name": "code",
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
        "/oauth/password": {
            "post": {
                "description": "PasswordOauthLogin",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Oauth"
                ],
                "parameters": [
                    {
                        "description": "PasswordRequest Model",
                        "name": "passwordRequest",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/oauth.PasswordRequest"
                        }
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
        "controller.ChangeStrategyStatusRequest": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                },
                "strategy_code": {
                    "type": "string"
                }
            }
        },
        "model.ApplyApiKeyRequest": {
            "type": "object",
            "properties": {
                "api_key_name": {
                    "type": "string"
                },
                "domain_whitelist": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "erc20_paymaster_enable": {
                    "type": "boolean"
                },
                "ip_white_list": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "network_limit_enable": {
                    "type": "boolean"
                },
                "paymaster_enable": {
                    "type": "boolean"
                },
                "project_sponsor_paymaster_enable": {
                    "type": "boolean"
                },
                "user_pay_paymaster_enable": {
                    "type": "boolean"
                }
            }
        },
        "model.UploadApiKeyRequest": {
            "type": "object",
            "properties": {
                "api_key": {
                    "type": "string"
                },
                "api_key_name": {
                    "type": "string"
                },
                "domain_whitelist": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "erc20_paymaster_enable": {
                    "type": "boolean"
                },
                "ip_white_list": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "network_limit_enable": {
                    "type": "boolean"
                },
                "paymaster_enable": {
                    "type": "boolean"
                },
                "project_sponsor_paymaster_enable": {
                    "type": "boolean"
                },
                "user_pay_paymaster_enable": {
                    "type": "boolean"
                }
            }
        },
        "model.UploadStrategyRequest": {
            "type": "object",
            "properties": {
                "address_block_list": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "chain_id_whitelist": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "day_max_usd": {
                    "type": "number"
                },
                "description": {
                    "type": "string"
                },
                "end_time": {
                    "type": "integer"
                },
                "global_max_usd": {
                    "type": "number"
                },
                "per_user_max_usd": {
                    "type": "number"
                },
                "project_code": {
                    "type": "string"
                },
                "start_time": {
                    "type": "integer"
                },
                "strategy_code": {
                    "type": "string"
                },
                "strategy_name": {
                    "type": "string"
                }
            }
        },
        "oauth.PasswordRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "JWT": {
            "description": "Type 'Bearer \\\u003cTOKEN\\\u003e' to correctly set the AccessToken",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "v1",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "AAStar BackEndDashBoard API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
