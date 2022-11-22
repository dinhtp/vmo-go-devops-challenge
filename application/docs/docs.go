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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/applications": {
            "get": {
                "description": "Get applications list by page and limit",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "applications"
                ],
                "summary": "Get applications list",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "current page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "current page",
                        "name": "type",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "page limit",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "description": "page limit",
                        "name": "enabled",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/message.ListApplicationResponse"
                        }
                    },
                    "500": {
                        "description": "{\"error\":\"error_code\", \"message\":\"error_description\"}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "post": {
                "description": "Create new application",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "applications"
                ],
                "summary": "Create new application",
                "parameters": [
                    {
                        "description": "Application Data",
                        "name": "application",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/message.Application"
                        }
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/message.Application"
                        }
                    },
                    "500": {
                        "description": "{\"error\":\"error_code\", \"message\":\"error_description\"}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        },
        "/applications/{id}": {
            "get": {
                "description": "Get an application by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "applications"
                ],
                "summary": "Get an application detail",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Application ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/message.Application"
                        }
                    },
                    "500": {
                        "description": "{\"error\":\"error_code\", \"message\":\"error_description\"}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "put": {
                "description": "Update application detail by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "applications"
                ],
                "summary": "Update application detail",
                "parameters": [
                    {
                        "description": "Application Data",
                        "name": "application",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/message.Application"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Application ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/message.Application"
                        }
                    },
                    "500": {
                        "description": "{\"error\":\"error_code\", \"message\":\"error_description\"}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete an application detail by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "applications"
                ],
                "summary": "Delete an application",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Application ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{}",
                        "schema": {
                            "type": "object"
                        }
                    },
                    "500": {
                        "description": "{\"error\":\"error_code\", \"message\":\"error_description\"}",
                        "schema": {
                            "type": "object"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "message.Application": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "enabled": {
                    "type": "boolean"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "message.ListApplicationResponse": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/message.Application"
                    }
                },
                "limit": {
                    "type": "integer"
                },
                "max_page": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                },
                "total_count": {
                    "type": "integer"
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
