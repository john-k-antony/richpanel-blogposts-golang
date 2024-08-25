// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "REST API implementation for a simple blog post system",
    "title": "BlogPost",
    "contact": {
      "name": "John Antony",
      "email": "john.k.antony@gmail.com"
    },
    "version": "1.0"
  },
  "host": "localhost:3000",
  "basePath": "/",
  "paths": {
    "/posts": {
      "get": {
        "security": [
          {
            "APIKey": []
          }
        ],
        "description": "Retrieve the information of the blog posts.",
        "produces": [
          "application/json"
        ],
        "tags": [
          "BlogPost"
        ],
        "summary": "List blog posts, with pagination",
        "operationId": "listBlogposts",
        "parameters": [
          {
            "type": "integer",
            "description": "The number of items to skip before starting to collect the result set",
            "name": "offset",
            "in": "query"
          },
          {
            "type": "integer",
            "description": "The numbers of items to return",
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Blog posts",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Blogpost"
              }
            },
            "headers": {
              "Pagination-Count": {
                "type": "integer",
                "description": "Total number of items"
              },
              "Pagination-Limit": {
                "type": "integer",
                "description": "Number of returned items"
              },
              "Pagination-Page": {
                "type": "integer",
                "description": "Current page number"
              }
            }
          },
          "401": {
            "description": "Unauthenticated Request",
            "schema": {
              "$ref": "#/definitions/ResponseError"
            }
          }
        },
        "x-eov-operation-handler": "handlers/blogpostsController"
      },
      "post": {
        "security": [
          {
            "APIKey": []
          }
        ],
        "description": "Create a new blog post.",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "BlogPost"
        ],
        "summary": "Create a new blog post",
        "operationId": "createBlogpost",
        "parameters": [
          {
            "description": "Post the necessary fields for the API to create a new blog post.",
            "name": "body",
            "in": "body",
            "schema": {
              "type": "object",
              "required": [
                "title"
              ],
              "properties": {
                "contents": {
                  "type": "string",
                  "x-nullable": true
                },
                "title": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Blog post created",
            "schema": {
              "$ref": "#/definitions/Blogpost"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/ResponseError"
            }
          },
          "401": {
            "description": "Unauthenticated Request",
            "schema": {
              "$ref": "#/definitions/ResponseError"
            }
          },
          "422": {
            "description": "Invalid request",
            "schema": {
              "$ref": "#/definitions/ResponseError"
            }
          }
        },
        "x-eov-operation-handler": "handlers/blogpostsController"
      }
    },
    "/posts/{id}": {
      "get": {
        "security": [
          {
            "APIKey": []
          }
        ],
        "description": "Retrieve the information of the blog post with the matching id.",
        "produces": [
          "application/json"
        ],
        "tags": [
          "BlogPost"
        ],
        "summary": "Get blog post by id",
        "operationId": "getBlogpostById",
        "parameters": [
          {
            "type": "string",
            "description": "Id of an existing blog post.",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Blog post found",
            "schema": {
              "$ref": "#/definitions/Blogpost"
            }
          },
          "401": {
            "description": "Unauthenticated Request",
            "schema": {
              "$ref": "#/definitions/ResponseError"
            }
          },
          "404": {
            "description": "Blog post not found",
            "schema": {
              "$ref": "#/definitions/ResponseError"
            }
          }
        },
        "x-eov-operation-handler": "handlers/blogpostsController"
      },
      "put": {
        "security": [
          {
            "APIKey": []
          }
        ],
        "description": "Update the information of an existing blog post with the matching id.",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "BlogPost"
        ],
        "summary": "Update blog post information",
        "operationId": "updateBlogpostById",
        "parameters": [
          {
            "type": "string",
            "description": "Id of an existing blog post.",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "description": "Put blog post properties to update.",
            "name": "body",
            "in": "body",
            "schema": {
              "type": "object",
              "properties": {
                "contents": {
                  "type": "string",
                  "x-nullable": true
                },
                "title": {
                  "type": "string",
                  "x-nullable": true
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Blog post updated",
            "schema": {
              "$ref": "#/definitions/Blogpost"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/ResponseError"
            }
          },
          "401": {
            "description": "Unauthenticated Request",
            "schema": {
              "$ref": "#/definitions/ResponseError"
            }
          },
          "404": {
            "description": "Blog post not found",
            "schema": {
              "$ref": "#/definitions/ResponseError"
            }
          },
          "422": {
            "description": "Invalid request",
            "schema": {
              "$ref": "#/definitions/ResponseError"
            }
          }
        },
        "x-eov-operation-handler": "handlers/blogpostsController"
      },
      "delete": {
        "security": [
          {
            "APIKey": []
          }
        ],
        "description": "Delete an existing blog post with the matching id.",
        "produces": [
          "application/json"
        ],
        "tags": [
          "BlogPost"
        ],
        "summary": "Delete blog post information",
        "operationId": "deleteBlogpostById",
        "parameters": [
          {
            "type": "string",
            "description": "Id of an existing blog post.",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Blog post deleted",
            "schema": {
              "$ref": "#/definitions/Blogpost"
            }
          },
          "401": {
            "description": "Unauthenticated Request",
            "schema": {
              "$ref": "#/definitions/ResponseError"
            }
          },
          "404": {
            "description": "Blog post not found",
            "schema": {
              "$ref": "#/definitions/ResponseError"
            }
          }
        },
        "x-eov-operation-handler": "handlers/blogpostsController"
      }
    }
  },
  "definitions": {
    "Blogpost": {
      "description": "Blog post object",
      "type": "object",
      "title": "Blog Post",
      "required": [
        "id",
        "title",
        "userId"
      ],
      "properties": {
        "contents": {
          "type": "string"
        },
        "createdAt": {
          "description": "The date that blog post was created.",
          "type": "string",
          "format": "date-time",
          "x-nullable": true
        },
        "id": {
          "description": "Unique identifier for the given blog post.",
          "type": "string"
        },
        "modifiedAt": {
          "description": "The date that blog post was modified.",
          "type": "string",
          "format": "date-time",
          "x-nullable": true
        },
        "title": {
          "type": "string"
        },
        "userId": {
          "description": "Id of the user that created the blog post",
          "type": "string"
        }
      }
    },
    "ResponseError": {
      "description": "Response error object",
      "type": "object",
      "title": "Response Error",
      "properties": {
        "errors": {
          "description": "Nested errors if any.",
          "type": "object",
          "x-nullable": true
        },
        "message": {
          "description": "Error message.",
          "type": "string",
          "x-nullable": true
        }
      }
    }
  },
  "securityDefinitions": {
    "APIKey": {
      "description": "JWT based authorization",
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  },
  "tags": [
    {
      "description": "BlogPost API",
      "name": "BlogPost"
    }
  ],
  "x-components": {}
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "REST API implementation for a simple blog post system",
    "title": "BlogPost",
    "contact": {
      "name": "John Antony",
      "email": "john.k.antony@gmail.com"
    },
    "version": "1.0"
  },
  "host": "localhost:3000",
  "basePath": "/",
  "paths": {
    "/posts": {
      "get": {
        "security": [
          {
            "APIKey": []
          }
        ],
        "description": "Retrieve the information of the blog posts.",
        "produces": [
          "application/json"
        ],
        "tags": [
          "BlogPost"
        ],
        "summary": "List blog posts, with pagination",
        "operationId": "listBlogposts",
        "parameters": [
          {
            "type": "integer",
            "description": "The number of items to skip before starting to collect the result set",
            "name": "offset",
            "in": "query"
          },
          {
            "type": "integer",
            "description": "The numbers of items to return",
            "name": "limit",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Blog posts",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/Blogpost"
              }
            },
            "headers": {
              "Pagination-Count": {
                "type": "integer",
                "description": "Total number of items"
              },
              "Pagination-Limit": {
                "type": "integer",
                "description": "Number of returned items"
              },
              "Pagination-Page": {
                "type": "integer",
                "description": "Current page number"
              }
            }
          },
          "401": {
            "description": "Unauthenticated Request",
            "schema": {
              "$ref": "#/definitions/ResponseError"
            }
          }
        },
        "x-eov-operation-handler": "handlers/blogpostsController"
      },
      "post": {
        "security": [
          {
            "APIKey": []
          }
        ],
        "description": "Create a new blog post.",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "BlogPost"
        ],
        "summary": "Create a new blog post",
        "operationId": "createBlogpost",
        "parameters": [
          {
            "description": "Post the necessary fields for the API to create a new blog post.",
            "name": "body",
            "in": "body",
            "schema": {
              "type": "object",
              "required": [
                "title"
              ],
              "properties": {
                "contents": {
                  "type": "string",
                  "x-nullable": true
                },
                "title": {
                  "type": "string"
                }
              }
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Blog post created",
            "schema": {
              "$ref": "#/definitions/Blogpost"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/ResponseError"
            }
          },
          "401": {
            "description": "Unauthenticated Request",
            "schema": {
              "$ref": "#/definitions/ResponseError"
            }
          },
          "422": {
            "description": "Invalid request",
            "schema": {
              "$ref": "#/definitions/ResponseError"
            }
          }
        },
        "x-eov-operation-handler": "handlers/blogpostsController"
      }
    },
    "/posts/{id}": {
      "get": {
        "security": [
          {
            "APIKey": []
          }
        ],
        "description": "Retrieve the information of the blog post with the matching id.",
        "produces": [
          "application/json"
        ],
        "tags": [
          "BlogPost"
        ],
        "summary": "Get blog post by id",
        "operationId": "getBlogpostById",
        "parameters": [
          {
            "type": "string",
            "description": "Id of an existing blog post.",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Blog post found",
            "schema": {
              "$ref": "#/definitions/Blogpost"
            }
          },
          "401": {
            "description": "Unauthenticated Request",
            "schema": {
              "$ref": "#/definitions/ResponseError"
            }
          },
          "404": {
            "description": "Blog post not found",
            "schema": {
              "$ref": "#/definitions/ResponseError"
            }
          }
        },
        "x-eov-operation-handler": "handlers/blogpostsController"
      },
      "put": {
        "security": [
          {
            "APIKey": []
          }
        ],
        "description": "Update the information of an existing blog post with the matching id.",
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "BlogPost"
        ],
        "summary": "Update blog post information",
        "operationId": "updateBlogpostById",
        "parameters": [
          {
            "type": "string",
            "description": "Id of an existing blog post.",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "description": "Put blog post properties to update.",
            "name": "body",
            "in": "body",
            "schema": {
              "type": "object",
              "properties": {
                "contents": {
                  "type": "string",
                  "x-nullable": true
                },
                "title": {
                  "type": "string",
                  "x-nullable": true
                }
              }
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Blog post updated",
            "schema": {
              "$ref": "#/definitions/Blogpost"
            }
          },
          "400": {
            "description": "Bad request",
            "schema": {
              "$ref": "#/definitions/ResponseError"
            }
          },
          "401": {
            "description": "Unauthenticated Request",
            "schema": {
              "$ref": "#/definitions/ResponseError"
            }
          },
          "404": {
            "description": "Blog post not found",
            "schema": {
              "$ref": "#/definitions/ResponseError"
            }
          },
          "422": {
            "description": "Invalid request",
            "schema": {
              "$ref": "#/definitions/ResponseError"
            }
          }
        },
        "x-eov-operation-handler": "handlers/blogpostsController"
      },
      "delete": {
        "security": [
          {
            "APIKey": []
          }
        ],
        "description": "Delete an existing blog post with the matching id.",
        "produces": [
          "application/json"
        ],
        "tags": [
          "BlogPost"
        ],
        "summary": "Delete blog post information",
        "operationId": "deleteBlogpostById",
        "parameters": [
          {
            "type": "string",
            "description": "Id of an existing blog post.",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Blog post deleted",
            "schema": {
              "$ref": "#/definitions/Blogpost"
            }
          },
          "401": {
            "description": "Unauthenticated Request",
            "schema": {
              "$ref": "#/definitions/ResponseError"
            }
          },
          "404": {
            "description": "Blog post not found",
            "schema": {
              "$ref": "#/definitions/ResponseError"
            }
          }
        },
        "x-eov-operation-handler": "handlers/blogpostsController"
      }
    }
  },
  "definitions": {
    "Blogpost": {
      "description": "Blog post object",
      "type": "object",
      "title": "Blog Post",
      "required": [
        "id",
        "title",
        "userId"
      ],
      "properties": {
        "contents": {
          "type": "string"
        },
        "createdAt": {
          "description": "The date that blog post was created.",
          "type": "string",
          "format": "date-time",
          "x-nullable": true
        },
        "id": {
          "description": "Unique identifier for the given blog post.",
          "type": "string"
        },
        "modifiedAt": {
          "description": "The date that blog post was modified.",
          "type": "string",
          "format": "date-time",
          "x-nullable": true
        },
        "title": {
          "type": "string"
        },
        "userId": {
          "description": "Id of the user that created the blog post",
          "type": "string"
        }
      }
    },
    "ResponseError": {
      "description": "Response error object",
      "type": "object",
      "title": "Response Error",
      "properties": {
        "errors": {
          "description": "Nested errors if any.",
          "type": "object",
          "x-nullable": true
        },
        "message": {
          "description": "Error message.",
          "type": "string",
          "x-nullable": true
        }
      }
    }
  },
  "securityDefinitions": {
    "APIKey": {
      "description": "JWT based authorization",
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  },
  "tags": [
    {
      "description": "BlogPost API",
      "name": "BlogPost"
    }
  ],
  "x-components": {}
}`))
}
