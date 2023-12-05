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
            "name": "Aamir Zaheer",
            "email": "aamirzaheer95@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/filterBooks": {
            "get": {
                "security": [
                    {
                        "bearerToken": []
                    }
                ],
                "description": "User can filter books author or publisher name also sort books by number of pages or average rating",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "FilterBooks user route",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Author name to filter by",
                        "name": "author",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Publisher name to filter by",
                        "name": "publisher",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Sort by average rating ASC or DESC",
                        "name": "avg_rating",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Sort by number of pages ASC or DESC",
                        "name": "num_pages",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.Books"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schema.Error"
                        }
                    }
                }
            }
        },
        "/getBooks": {
            "get": {
                "security": [
                    {
                        "bearerToken": []
                    }
                ],
                "description": "User can get the list of books with pagination",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "GetBooks user route",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page number (default 1)",
                        "name": "page_no",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Number of itmes per page (default 10)",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.Books"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/schema.Error"
                        }
                    }
                }
            }
        },
        "/giveBookBack": {
            "put": {
                "security": [
                    {
                        "bearerToken": []
                    }
                ],
                "description": "User can give the rented book back to the admin and admin can update the user rent details",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "Give Book back route",
                "parameters": [
                    {
                        "description": "Request body in JSON format",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.GiveBookBackDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.RentBookSuccess"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/schema.Error"
                        }
                    }
                }
            }
        },
        "/rentBook/{id}": {
            "post": {
                "security": [
                    {
                        "bearerToken": []
                    }
                ],
                "description": "User can rent a book for a period of time",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "books"
                ],
                "summary": "RentBook route",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the book to rent",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Request body in JSON format",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schema.RentBookDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/schema.RentBookSuccess"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/schema.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "schema.Books": {
            "type": "object",
            "properties": {
                "authors": {
                    "type": "string"
                },
                "avg_rating": {
                    "type": "number"
                },
                "language_code": {
                    "type": "string"
                },
                "num_pages": {
                    "type": "integer"
                },
                "publication_date": {
                    "type": "string"
                },
                "publisher": {
                    "type": "string"
                },
                "text_reviews_count": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "schema.Error": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "statuscode": {
                    "type": "integer"
                },
                "statustext": {
                    "type": "string"
                }
            }
        },
        "schema.GiveBookBackDTO": {
            "type": "object",
            "properties": {
                "book_id": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                }
            }
        },
        "schema.RentBookDTO": {
            "type": "object",
            "properties": {
                "rentDuration": {
                    "type": "object",
                    "properties": {
                        "days": {
                            "type": "integer"
                        },
                        "months": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "schema.RentBookSuccess": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status_code": {
                    "type": "integer"
                },
                "status_text": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "bearerToken": {
            "description": "Enter your access_token in the form of \u003cb\u003eBearer \u0026lt;access_token\u0026gt;\u003c/b\u003e",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Books Api",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
