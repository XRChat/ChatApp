// Package swagger Code generated by swaggo/swag. DO NOT EDIT
package swagger

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "https://example.com/terms",
        "contact": {
            "name": "Mohammad Developer",
            "url": "https://example.com",
            "email": "mohammad.developer@example.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/user": {
            "post": {
                "description": "Create a new user with the provided details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "User registration details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/CreateUserInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User registered successfully",
                        "schema": {
                            "$ref": "#/definitions/User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/InvalidRequestError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/UnauthorizedError"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/ForbiddenAccessError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/SystemError"
                        }
                    }
                }
            }
        },
        "/user/{id}": {
            "get": {
                "security": [
                    {
                        "JWT": []
                    }
                ],
                "description": "Find a user based on the provided ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Find a user by ID",
                "operationId": "findUserByID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Bearer ",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/BaseResponse"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/User"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/InvalidRequestError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/UnauthorizedError"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/ForbiddenAccessError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/SystemError"
                        }
                    }
                }
            }
        },
        "/user/{username}": {
            "get": {
                "description": "Get user information by their username",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Find a user by username",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Username",
                        "name": "username",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User details",
                        "schema": {
                            "$ref": "#/definitions/User"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/InvalidRequestError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/UnauthorizedError"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/ForbiddenAccessError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/SystemError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "BaseResponse": {
            "type": "object",
            "properties": {
                "data": {}
            }
        },
        "CreateUserInput": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string",
                    "example": "password123"
                },
                "role": {
                    "allOf": [
                        {
                            "$ref": "#/definitions/github_com_chatApp_internal_domain.UserRole"
                        }
                    ],
                    "example": "ADMIN"
                },
                "user_name": {
                    "type": "string",
                    "example": "+919984778491"
                }
            }
        },
        "ForbiddenAccessError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string",
                    "example": "FORBIDDEN_ACCESS"
                },
                "message": {
                    "type": "string",
                    "example": "You are forbidden from accessing this resource"
                }
            }
        },
        "InvalidRequestError": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "invalid request"
                }
            }
        },
        "SystemError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string",
                    "example": "INTERNAL_SERVER_ERROR"
                },
                "message": {
                    "type": "string",
                    "example": "Oops! Something went wrong. Please try again later"
                }
            }
        },
        "UnauthorizedError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string",
                    "example": "UNAUTHORIZED"
                },
                "message": {
                    "type": "string",
                    "example": "You are not authorized to access this resource"
                }
            }
        },
        "User": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string",
                    "example": ""
                },
                "password": {
                    "type": "string"
                },
                "role": {
                    "type": "string",
                    "example": "ADMIN"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_name": {
                    "type": "string",
                    "example": "+919984778491"
                }
            }
        },
        "github_com_chatApp_internal_domain.UserRole": {
            "type": "string",
            "enum": [
                "ADMIN",
                "USER"
            ],
            "x-enum-varnames": [
                "UserRoleAmin",
                "UserRoleUser"
            ]
        }
    },
    "securityDefinitions": {
        "JWT": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{"http", "https"},
	Title:            "Chat API",
	Description:      "Chat application's set of APIs",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}