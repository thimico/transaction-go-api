// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Thiago Menezes",
            "email": "thg.mnzs@gmail.com"
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
		"/health": {
			"get": {
                "description": "Get health status",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Health"
                ],
                "summary": "Get health status",
                "responses": {
                    "200": {
                        "description": "OK",
                    },
					"500": {
						"description": "NOT OK",
					}
                }
            },
		},
        "/accounts": {
            "post": {
                "description": "Create a new account with the input payload",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Accounts"
                ],
                "summary": "Create a new account",
                "parameters": [
                    {
                        "description": "Create account",
                        "name": "account",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AccountIn"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Account created",
                        "schema": {
                            "$ref": "#/definitions/model.AccountOut"
                        }
                    }
                }
            }
        },
        "/accounts/{id}": {
            "get": {
                "description": "Get account details by ID",
                "summary": "Get account by ID",
                "tags": [
                    "Accounts"
                ],
                "produces": [
                    "application/json"
                ],
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "description": "ID of the account",
                        "required": true,
                        "type": "string"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "successful operation",
                        "schema": {
                            "$ref": "#/definitions/model.AccountOut"
                        }
                    },
                    "404": {
                        "description": "Account not found"
                    }
                }
            }
        },
        "/transactions": {
            "post": {
                "description": "Create a new transaction with the input payload",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transactions"
                ],
                "summary": "Create a new transaction",
                "parameters": [
                    {
                        "description": "Create transaction",
                        "name": "transaction",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.TransactionIn"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Transaction created",
                        "schema": {
                            "$ref": "#/definitions/model.TransactionOut"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.AccountIn": {
            "type": "object",
            "required": [
                "document_number"
            ],
            "properties": {
                "document_number": {
                    "type": "string",
                    "description": "Document number of the account",
                    "example": "12345678900"
                }
            }
        },
        "model.AccountOut": {
            "type": "object",
            "properties": {
                "account_id": {
                    "type": "integer",
                    "description": "ID of the account"
                },
                "document_number": {
                    "type": "string",
                    "description": "Document number of the account",
                    "example": "12345678900"
                }
            }
        },
        "model.TransactionIn": {
            "type": "object",
            "required": [
                "account_id",
                "operation_type_id",
                "amount"
            ],
            "properties": {
                "account_id": {
                    "type": "integer",
                    "description": "ID of the account"
                },
                "operation_type_id": {
                    "type": "integer",
                    "description": "Type of the operation",
                    "example": 4
                },
                "amount": {
                    "type": "number",
                    "format": "float",
                    "description": "Amount of the transaction",
                    "example": 123.45
                }
            }
        },
        "model.TransactionOut": {
            "type": "object",
            "properties": {
                "transaction_id": {
                    "type": "integer",
                    "description": "ID of the transaction"
                },
                "account_id": {
                    "type": "integer",
                    "description": "ID of the account"
                },
                "operation_type_id": {
                    "type": "integer",
                    "description": "Type of the operation",
                    "example": 4
                },
                "amount": {
                    "type": "number",
                    "format": "float",
                    "description": "Amount of the transaction",
                    "example": 123.45"
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
	BasePath:    "/",
	Schemes:     []string{},
	Title:       "Pismo Challenge API",
	Description: "Pismo Challenge API",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

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
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}