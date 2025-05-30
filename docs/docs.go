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
        "/api": {
            "get": {
                "description": "Get information about users with filtering options",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get users information",
                "operationId": "get-users-info",
                "parameters": [
                    {
                        "minimum": 1,
                        "type": "integer",
                        "example": 10,
                        "description": "Limit number of records",
                        "name": "size",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "\"Ivan\"",
                        "description": "Filter by name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "\"Ivanov\"",
                        "description": "Filter by surname",
                        "name": "surname",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "example": "\"Ivanovich\"",
                        "description": "Filter by patronymic",
                        "name": "patronymic",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Slice human structs with all fields",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/github_com_Wladim1r_testtask_internal_models.Human"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid request data",
                        "schema": {
                            "$ref": "#/definitions/github_com_Wladim1r_testtask_internal_models.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Record not found",
                        "schema": {
                            "$ref": "#/definitions/github_com_Wladim1r_testtask_internal_models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Database error",
                        "schema": {
                            "$ref": "#/definitions/github_com_Wladim1r_testtask_internal_models.ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create user full name and automatically add age, gender and nationality",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Create new user",
                "operationId": "post-user",
                "parameters": [
                    {
                        "description": "User data to create",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_Wladim1r_testtask_internal_models.PostRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "User created successfully",
                        "schema": {
                            "$ref": "#/definitions/github_com_Wladim1r_testtask_internal_models.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request data",
                        "schema": {
                            "$ref": "#/definitions/github_com_Wladim1r_testtask_internal_models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error (API failure or database error)",
                        "schema": {
                            "$ref": "#/definitions/github_com_Wladim1r_testtask_internal_models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/{id}": {
            "delete": {
                "description": "Permanently removes user information by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Delete user information",
                "operationId": "delete-user",
                "parameters": [
                    {
                        "minimum": 1,
                        "type": "integer",
                        "example": 3,
                        "description": "ID of the user to delete",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User deleted successfully",
                        "schema": {
                            "$ref": "#/definitions/github_com_Wladim1r_testtask_internal_models.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request data",
                        "schema": {
                            "$ref": "#/definitions/github_com_Wladim1r_testtask_internal_models.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Record not found",
                        "schema": {
                            "$ref": "#/definitions/github_com_Wladim1r_testtask_internal_models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Database or Internal Server error",
                        "schema": {
                            "$ref": "#/definitions/github_com_Wladim1r_testtask_internal_models.ErrorResponse"
                        }
                    }
                }
            },
            "patch": {
                "description": "Change fields that were transmitted",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Update user information",
                "operationId": "change-user",
                "parameters": [
                    {
                        "minimum": 1,
                        "type": "integer",
                        "example": 13,
                        "description": "ID of the user to change",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "New data for change existing data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_Wladim1r_testtask_internal_models.Human"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User changed successfully",
                        "schema": {
                            "$ref": "#/definitions/github_com_Wladim1r_testtask_internal_models.SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request data",
                        "schema": {
                            "$ref": "#/definitions/github_com_Wladim1r_testtask_internal_models.ErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Record not found",
                        "schema": {
                            "$ref": "#/definitions/github_com_Wladim1r_testtask_internal_models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/github_com_Wladim1r_testtask_internal_models.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "github_com_Wladim1r_testtask_internal_models.ErrorResponse": {
            "description": "Default error response",
            "type": "object",
            "properties": {
                "error": {
                    "type": "string",
                    "example": "error description"
                }
            }
        },
        "github_com_Wladim1r_testtask_internal_models.Human": {
            "description": "Detailed information about a person",
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer",
                    "example": 18
                },
                "gender": {
                    "type": "string",
                    "example": "male"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "Vladimir"
                },
                "nationality": {
                    "type": "string",
                    "example": "RU"
                },
                "patronymic": {
                    "type": "string",
                    "example": "Dmitrievich"
                },
                "surname": {
                    "type": "string",
                    "example": "Sokolov"
                }
            }
        },
        "github_com_Wladim1r_testtask_internal_models.PostRequest": {
            "description": "Required information for creating a new user",
            "type": "object",
            "properties": {
                "name": {
                    "type": "string",
                    "example": "Vladimir"
                },
                "patronymic": {
                    "type": "string",
                    "example": "Dmitrievich"
                },
                "surname": {
                    "type": "string",
                    "example": "Sokolov"
                }
            }
        },
        "github_com_Wladim1r_testtask_internal_models.SuccessResponse": {
            "description": "Default successfully response",
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "message description"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.3",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "API Server",
	Description:      "API for managing human information with automatic age, gender and nationality detection",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
