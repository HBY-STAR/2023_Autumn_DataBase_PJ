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
            "name": "Maintainer Polaris",
            "email": "polaris@fduhole.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "https://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/commodity/all": {
            "get": {
                "description": "获取所有商品",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Item"
                ],
                "summary": "获取所有商品",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.CommodityItem"
                        }
                    }
                }
            }
        },
        "/api/commodity/item": {
            "put": {
                "description": "更新商品",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Item"
                ],
                "summary": "更新商品",
                "parameters": [
                    {
                        "description": "json",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/item.UpdateItemModel"
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
                "description": "添加商品",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Item"
                ],
                "summary": "添加商品",
                "parameters": [
                    {
                        "description": "json",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/item.CreateItemModel"
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
                "description": "删除商品",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Item"
                ],
                "summary": "删除商品",
                "parameters": [
                    {
                        "type": "string",
                        "description": "commodity id",
                        "name": "id",
                        "in": "path",
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
        "/api/favorites": {
            "post": {
                "description": "AddFavorite",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Favorite"
                ],
                "summary": "AddFavorite",
                "parameters": [
                    {
                        "description": "json",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/favorite.AddFavoriteModel"
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
        "/api/favorites/all": {
            "get": {
                "description": "GetAllFavorites",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Favorite"
                ],
                "summary": "GetAllFavorites",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Favorite"
                        }
                    }
                }
            }
        },
        "/api/message/all": {
            "get": {
                "description": "Get all messages",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Message"
                ],
                "summary": "Get all messages",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Message"
                        }
                    }
                }
            }
        },
        "/api/platform": {
            "put": {
                "description": "Update platform",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Platform"
                ],
                "summary": "Update platform",
                "parameters": [
                    {
                        "description": "json",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/platform.UpdatePlatformRequest"
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
                "description": "Add platform",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Platform"
                ],
                "summary": "Add platform",
                "parameters": [
                    {
                        "description": "json",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/platform.CreatePlatformRequest"
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
                "description": "Delete platform",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Platform"
                ],
                "summary": "Delete platform",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "platform id",
                        "name": "id",
                        "in": "path",
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
        "/api/platform/all": {
            "get": {
                "description": "Get all platform",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Platform"
                ],
                "summary": "Get all platform",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Platform"
                        }
                    }
                }
            }
        },
        "/api/price/limit": {
            "post": {
                "description": "AddPriceLimit",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Favorite"
                ],
                "summary": "AddPriceLimit",
                "parameters": [
                    {
                        "description": "json",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/favorite.PriceLimitModel"
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
        "/api/search": {
            "post": {
                "description": "搜索商品",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Item"
                ],
                "summary": "搜索商品",
                "parameters": [
                    {
                        "description": "Range: name, category; Search: content to search",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/item.SearchQuery"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.CommodityItem"
                        }
                    }
                }
            }
        },
        "/api/sellers": {
            "put": {
                "description": "Update seller",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Seller"
                ],
                "summary": "Update seller",
                "parameters": [
                    {
                        "description": "json",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/seller.UpdateSellerRequest"
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
                "description": "Add seller",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Seller"
                ],
                "summary": "Add seller",
                "parameters": [
                    {
                        "description": "json",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/seller.CreateSellerRequest"
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
                "description": "Delete seller",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Seller"
                ],
                "summary": "Delete seller",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "seller id",
                        "name": "id",
                        "in": "path",
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
        "/api/sellers/data": {
            "get": {
                "description": "Get seller info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Seller"
                ],
                "summary": "Get seller info",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Seller"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/common.HttpError"
                        }
                    }
                }
            }
        },
        "/api/users": {
            "put": {
                "description": "Update user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Update user",
                "parameters": [
                    {
                        "description": "json",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.UpdateUserRequest"
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
                "description": "Add user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Add user",
                "parameters": [
                    {
                        "description": "json",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.CreateUserRequest"
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
                "description": "Delete user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Delete user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "user id",
                        "name": "id",
                        "in": "path",
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
        "/api/users/data": {
            "get": {
                "description": "Get user info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Get user info",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.User"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/common.HttpError"
                        }
                    }
                }
            }
        },
        "/commodities": {
            "put": {
                "description": "更新商品",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "商品"
                ],
                "summary": "更新商品",
                "parameters": [
                    {
                        "description": "json",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.Commodity"
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
                "description": "添加商品",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "商品"
                ],
                "summary": "添加商品",
                "parameters": [
                    {
                        "description": "json",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/commodity.CreateCommodityRequest"
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
                "description": "删除商品",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "商品"
                ],
                "summary": "删除商品",
                "parameters": [
                    {
                        "type": "string",
                        "description": "commodity id",
                        "name": "id",
                        "in": "path",
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
        "/commodities/all": {
            "get": {
                "description": "获取所有商品",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "商品"
                ],
                "summary": "获取所有商品",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Commodity"
                        }
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Login",
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
                        "description": "json",
                        "name": "json",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/auth.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/auth.TokenResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.HttpError"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/common.HttpError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/common.HttpError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "auth.LoginRequest": {
            "type": "object",
            "required": [
                "password",
                "type",
                "username"
            ],
            "properties": {
                "password": {
                    "type": "string",
                    "minLength": 3
                },
                "type": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "auth.TokenResponse": {
            "type": "object",
            "properties": {
                "access": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "commodity.CreateCommodityRequest": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "default_name": {
                    "type": "string"
                },
                "produce_address": {
                    "type": "string"
                },
                "produce_at": {
                    "type": "string"
                }
            }
        },
        "common.ErrorDetailElement": {
            "type": "object",
            "properties": {
                "field": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "param": {
                    "type": "string"
                },
                "struct_field": {
                    "type": "string"
                },
                "tag": {
                    "type": "string"
                },
                "value": {}
            }
        },
        "common.HttpError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "detail": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/common.ErrorDetailElement"
                    }
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "favorite.AddFavoriteModel": {
            "type": "object",
            "required": [
                "item_id"
            ],
            "properties": {
                "item_id": {
                    "type": "integer"
                }
            }
        },
        "favorite.PriceLimitModel": {
            "type": "object",
            "properties": {
                "item_id": {
                    "type": "integer"
                },
                "price_limit": {
                    "type": "integer"
                }
            }
        },
        "item.CreateItemModel": {
            "type": "object",
            "required": [
                "commodity_id",
                "item_name",
                "platform_id",
                "price"
            ],
            "properties": {
                "commodity_id": {
                    "type": "integer"
                },
                "item_name": {
                    "type": "string"
                },
                "platform_id": {
                    "type": "integer"
                },
                "price": {
                    "type": "number"
                }
            }
        },
        "item.SearchQuery": {
            "type": "object",
            "properties": {
                "accurate": {
                    "type": "boolean"
                },
                "range": {
                    "type": "string"
                },
                "search": {
                    "type": "string"
                }
            }
        },
        "item.UpdateItemModel": {
            "type": "object",
            "required": [
                "commodity_item_id",
                "item_name",
                "price"
            ],
            "properties": {
                "commodity_item_id": {
                    "type": "integer"
                },
                "item_name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                }
            }
        },
        "models.Commodity": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "default_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "produce_address": {
                    "type": "string"
                },
                "produce_at": {
                    "type": "string"
                }
            }
        },
        "models.CommodityItem": {
            "type": "object",
            "properties": {
                "commodity": {
                    "$ref": "#/definitions/models.Commodity"
                },
                "commodity_id": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "item_name": {
                    "type": "string"
                },
                "platform": {
                    "$ref": "#/definitions/models.Platform"
                },
                "platform_id": {
                    "type": "integer"
                },
                "price": {
                    "type": "number"
                },
                "seller": {
                    "$ref": "#/definitions/models.Seller"
                },
                "seller_id": {
                    "type": "integer"
                },
                "update_at": {
                    "type": "string"
                }
            }
        },
        "models.Favorite": {
            "type": "object",
            "properties": {
                "commodityItem": {
                    "$ref": "#/definitions/models.CommodityItem"
                },
                "commodity_item_id": {
                    "type": "integer"
                },
                "create_at": {
                    "type": "string"
                },
                "price_limit": {
                    "type": "number"
                },
                "user": {
                    "$ref": "#/definitions/models.User"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "models.Message": {
            "type": "object",
            "properties": {
                "commodityItem": {
                    "$ref": "#/definitions/models.CommodityItem"
                },
                "commodity_item_id": {
                    "type": "integer"
                },
                "create_at": {
                    "type": "string"
                },
                "current_price": {
                    "type": "number"
                },
                "id": {
                    "type": "integer"
                },
                "user": {
                    "$ref": "#/definitions/models.User"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "models.Platform": {
            "type": "object",
            "properties": {
                "country": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "models.PriceChange": {
            "type": "object",
            "properties": {
                "commodityItem": {
                    "$ref": "#/definitions/models.CommodityItem"
                },
                "commodity_item_id": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "new_price": {
                    "type": "number"
                },
                "update_at": {
                    "type": "string"
                }
            }
        },
        "models.Seller": {
            "type": "object",
            "properties": {
                "address": {
                    "description": "Age      int",
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "gender": {
                    "type": "boolean"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "platform.CreatePlatformRequest": {
            "type": "object",
            "properties": {
                "country": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "platform.UpdatePlatformRequest": {
            "type": "object",
            "properties": {
                "country": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "priceChange.GetPriceChangeModel": {
            "type": "object",
            "required": [
                "commodity_item_id",
                "time_end",
                "time_start"
            ],
            "properties": {
                "commodity_item_id": {
                    "type": "integer"
                },
                "time_end": {
                    "type": "string"
                },
                "time_start": {
                    "type": "string"
                }
            }
        },
        "priceChange.UpdatePriceChangeModel": {
            "type": "object",
            "properties": {
                "commodity_item_id": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "new_price": {
                    "type": "number"
                }
            }
        },
        "seller.CreateSellerRequest": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "seller.UpdateSellerRequest": {
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
                "password": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "user.CreateUserRequest": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "gender": {
                    "type": "boolean"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "user.UpdateUserRequest": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "gender": {
                    "type": "boolean"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
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
	Version:          "1.0.0",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Database System Design Project",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
