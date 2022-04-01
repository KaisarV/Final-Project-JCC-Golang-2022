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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/cart": {
            "get": {
                "description": "display all cart items of users who are currently logged in.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cart"
                ],
                "summary": "Get all cart items.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.CartsResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "insert product to cart belongs to the user who is currently logged in.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cart"
                ],
                "summary": "insert cart.",
                "parameters": [
                    {
                        "description": "cart's data",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Cart"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.CartResponse"
                        }
                    }
                }
            }
        },
        "/cart/{cartId}": {
            "post": {
                "description": "update cart belongs to the user who is currently logged in.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cart"
                ],
                "summary": "update cart.",
                "parameters": [
                    {
                        "description": "cart's data",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Cart"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete display all cart items of users who are currently logged in.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cart"
                ],
                "summary": "delete cart item.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "cartId",
                        "name": "cartId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/feedbacks": {
            "get": {
                "description": "get a list of feedback from logged in users.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Feedbacks"
                ],
                "summary": "Get user feedback.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.FeedbacksResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "added feedback about the app so admin can see it.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Feedbacks"
                ],
                "summary": "insert feedback.",
                "parameters": [
                    {
                        "description": "feedback's data",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Feedback"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.FeedbackResponse"
                        }
                    }
                }
            }
        },
        "/feedbacks/all": {
            "get": {
                "description": "get a list of feedback from all users, only admin can use it.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Feedbacks"
                ],
                "summary": "Get all user feedback.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.FeedbacksResponse"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "login for registered users.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "login user.",
                "parameters": [
                    {
                        "description": "User's login data",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.InputLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/logout": {
            "get": {
                "description": "logout user.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "logout user.",
                "responses": {}
            }
        },
        "/product": {
            "post": {
                "description": "insert products sold by logged in users.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "insert product.",
                "parameters": [
                    {
                        "description": "product's data",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Product"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ProductResponse"
                        }
                    }
                }
            }
        },
        "/product/{productid}": {
            "put": {
                "description": "product updates sold by logged in users.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "update product's data.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "productid",
                        "name": "productid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "product's data",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Product"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ProductResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete products sold by logged in users.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "delete prodduct.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "productid",
                        "name": "productid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/products": {
            "get": {
                "description": "display all products.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Get all product.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ProductsResponse"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "insert user and it use for register user.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "insert user.",
                "parameters": [
                    {
                        "description": "User's data",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.InputUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.UserResponse"
                        }
                    }
                }
            }
        },
        "/review/{productid}": {
            "put": {
                "description": "update a review on the product that has been purchased.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Reviews"
                ],
                "summary": "update product's review.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "productid",
                        "name": "productid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "transaction's data",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ProductReview"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ProductReviewResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "leave a review on the product that has been purchased.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Reviews"
                ],
                "summary": "delete prodduct review.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "productid",
                        "name": "productid",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "review's data",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ProductReview"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ProductReviewResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete the review that has been given.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Reviews"
                ],
                "summary": "delete prodduct review.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "productid",
                        "name": "productid",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/reviews": {
            "get": {
                "description": "displays customer reviews given to things that have been purchased.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Reviews"
                ],
                "summary": "Get all product reviews.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ProductReviewsResponse"
                        }
                    }
                }
            }
        },
        "/store": {
            "put": {
                "description": "update the store of the currently logged in user.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Stores"
                ],
                "summary": "update store.",
                "parameters": [
                    {
                        "description": "store's data",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Store"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.StoreResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "added the store of the currently logged in user.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Stores"
                ],
                "summary": "insert store.",
                "parameters": [
                    {
                        "description": "store's data",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Store"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.StoreResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "delete the store belonging to the user who is currently logged in.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Stores"
                ],
                "summary": "delete store.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/stores": {
            "get": {
                "description": "Display all stores.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Stores"
                ],
                "summary": "Get all stores.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.StoresResponse"
                        }
                    }
                }
            }
        },
        "/transactions": {
            "get": {
                "description": "display all transactions of users who are currently logged in.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transactions"
                ],
                "summary": "Get all transactions.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.TransactionsResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "insert user's transaction who currently logged in.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transactions"
                ],
                "summary": "insert transaction.",
                "parameters": [
                    {
                        "description": "transaction's data",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Transaction"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.TransactionResponse"
                        }
                    }
                }
            }
        },
        "/user": {
            "put": {
                "description": "change the data of the user who is currently logged in.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "update user.",
                "parameters": [
                    {
                        "description": "User's data",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controllers.InputUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.UserResponse"
                        }
                    }
                }
            }
        },
        "/user/{id}": {
            "delete": {
                "description": "Delete user by id and admin only can use it.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "delete user.",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "description": "Display all registered users.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Get all users.",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.UsersResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.InputLogin": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "controllers.InputUser": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "model.Cart": {
            "type": "object",
            "properties": {
                "ProductId": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "qty": {
                    "type": "integer"
                },
                "userid": {
                    "type": "integer"
                }
            }
        },
        "model.CartResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/model.Cart"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "model.CartsResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Cart"
                    }
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "model.ErrorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "model.Feedback": {
            "type": "object",
            "properties": {
                "Date": {
                    "type": "string"
                },
                "feedback": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "userid": {
                    "type": "integer"
                }
            }
        },
        "model.FeedbackResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/model.Feedback"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "model.FeedbacksResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Feedback"
                    }
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "model.Product": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "storeId": {
                    "type": "integer"
                }
            }
        },
        "model.ProductResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/model.Product"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "model.ProductReview": {
            "type": "object",
            "properties": {
                "Date": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "productid": {
                    "type": "integer"
                },
                "rating": {
                    "type": "integer"
                },
                "review": {
                    "type": "string"
                },
                "userid": {
                    "type": "integer"
                }
            }
        },
        "model.ProductReviewResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/model.ProductReview"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "model.ProductReviewsResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.ProductReview"
                    }
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "model.ProductsResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Product"
                    }
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "model.Store": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "model.StoreResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/model.Store"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "model.StoresResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Store"
                    }
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "model.Transaction": {
            "type": "object",
            "properties": {
                "ProductId": {
                    "type": "integer"
                },
                "date": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "qty": {
                    "type": "integer"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "model.TransactionResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/model.Transaction"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "model.TransactionsResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Transaction"
                    }
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "usertype": {
                    "type": "integer"
                }
            }
        },
        "model.UserResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/model.User"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "model.UsersResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.User"
                    }
                },
                "message": {
                    "type": "string"
                },
                "status": {
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
