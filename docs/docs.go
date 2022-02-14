// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support (Thiago Mota)",
            "url": "http://engrenelog.com/es/",
            "email": "thiago@engrenelog.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/customers/paginated": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customers"
                ],
                "summary": "Lista de Clientes",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/customer.getCustomersRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/customer.CustomerList"
                        }
                    }
                }
            }
        },
        "/employees/": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Employees"
                ],
                "summary": "Actualizar Empleado",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/employee.updateEmployeeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Employees"
                ],
                "summary": "Insertar Empleado",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/employee.addEmployeeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/employees/best": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Employees"
                ],
                "summary": "Mejor Empleado",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/employee.BestEmployee"
                        }
                    }
                }
            }
        },
        "/employees/paginated": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Employees"
                ],
                "summary": "Lista de Empleados",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/employee.getEmployeesRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/employee.EmployeeList"
                        }
                    }
                }
            }
        },
        "/employees/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Employees"
                ],
                "summary": "Empleado By Id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Employee Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/employee.Employee"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Employees"
                ],
                "summary": "Eliminar Empleado",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Employee ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/orders/": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Actualizar Orden",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/order.addOrderRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Insertar Orden",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/order.addOrderRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/orders/paginated": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Lista de Ordenes",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/order.getOrdersRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/order.OrderList"
                        }
                    }
                }
            }
        },
        "/orders/{id}": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Eliminar Orden",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Order ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/orders/{orderId}/detail/{orderDetailId}": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Eliminar Detal de la Orden",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Order Detail ID",
                        "name": "orderDetailId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/orders{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Orders"
                ],
                "summary": "Order By Id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Order Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/order.OrderItem"
                        }
                    }
                }
            }
        },
        "/products/": {
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Actualizar Producto",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/product.updateProductRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Insertar Producto",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/product.getAddProductRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/products/bestSellers": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Mejores Productos",
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/product.ProductTop"
                        }
                    }
                }
            }
        },
        "/products/paginated": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Lista de Productos paginada",
                "parameters": [
                    {
                        "description": "User Data",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/product.getProductsRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "$ref": "#/definitions/product.ProductList"
                        }
                    }
                }
            }
        },
        "/products/{id}": {
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Eliminar Producto",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Employee ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "customer.Customer": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "businessphone": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "company": {
                    "type": "string"
                },
                "firstname": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "lastname": {
                    "type": "string"
                }
            }
        },
        "customer.CustomerList": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/customer.Customer"
                    }
                },
                "totalRecords": {
                    "type": "integer"
                }
            }
        },
        "customer.getCustomersRequest": {
            "type": "object",
            "properties": {
                "limit": {
                    "type": "integer"
                },
                "offset": {
                    "type": "integer"
                }
            }
        },
        "employee.BestEmployee": {
            "type": "object",
            "properties": {
                "Id": {
                    "type": "integer"
                },
                "firstName": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "totalVentas": {
                    "type": "integer"
                }
            }
        },
        "employee.Employee": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "businessPhone": {
                    "type": "string"
                },
                "company": {
                    "type": "string"
                },
                "emailAddress": {
                    "type": "string"
                },
                "faxNumber": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "homePhone": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "jobTitle": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "mobilePhone": {
                    "type": "string"
                }
            }
        },
        "employee.EmployeeList": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/employee.Employee"
                    }
                },
                "totalRecords": {
                    "type": "integer"
                }
            }
        },
        "employee.addEmployeeRequest": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "business": {
                    "type": "string"
                },
                "businessPhone": {
                    "type": "string"
                },
                "company": {
                    "type": "string"
                },
                "emailAddress": {
                    "type": "string"
                },
                "faxNumber": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "homePhone": {
                    "type": "string"
                },
                "jobTitle": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "mobilePhone": {
                    "type": "string"
                }
            }
        },
        "employee.getEmployeesRequest": {
            "type": "object",
            "properties": {
                "limit": {
                    "type": "integer"
                },
                "offset": {
                    "type": "integer"
                }
            }
        },
        "employee.updateEmployeeRequest": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "business": {
                    "type": "string"
                },
                "businessPhone": {
                    "type": "string"
                },
                "company": {
                    "type": "string"
                },
                "emailAddress": {
                    "type": "string"
                },
                "faxNumber": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "homePhone": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "jobTitle": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "mobilePhone": {
                    "type": "string"
                }
            }
        },
        "order.OrderDetailItem": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "order_id": {
                    "type": "integer"
                },
                "product_id": {
                    "type": "integer"
                },
                "product_name": {
                    "type": "string"
                },
                "quantity": {
                    "type": "number"
                },
                "unit_price": {
                    "type": "number"
                }
            }
        },
        "order.OrderItem": {
            "type": "object",
            "properties": {
                "address": {
                    "type": "string"
                },
                "city": {
                    "type": "string"
                },
                "company": {
                    "type": "string"
                },
                "customer": {
                    "type": "string"
                },
                "customerId": {
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/order.OrderDetailItem"
                    }
                },
                "orderDate": {
                    "type": "string"
                },
                "orderID": {
                    "type": "integer"
                },
                "orderName": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "statusId": {
                    "type": "string"
                }
            }
        },
        "order.OrderList": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/order.OrderItem"
                    }
                },
                "totlRecords": {
                    "type": "integer"
                }
            }
        },
        "order.addOrderDetailRequest": {
            "type": "object",
            "properties": {
                "id": {
                    "description": "Vamos utilizalas para el update acrescentando la variable ID",
                    "type": "integer"
                },
                "orderId": {
                    "type": "integer"
                },
                "productId": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                },
                "unitPrice": {
                    "type": "number"
                }
            }
        },
        "order.addOrderRequest": {
            "type": "object",
            "properties": {
                "customerId": {
                    "type": "integer"
                },
                "id": {
                    "description": "Vamos utilizalas para el update acrescentando la variable ID",
                    "type": "integer"
                },
                "orderDate": {
                    "type": "string"
                },
                "orderDetails": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/order.addOrderDetailRequest"
                    }
                }
            }
        },
        "order.getOrdersRequest": {
            "type": "object",
            "properties": {
                "dateFrom": {},
                "dateTo": {},
                "limit": {
                    "type": "integer"
                },
                "offset": {
                    "type": "integer"
                },
                "status": {}
            }
        },
        "product.Product": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "listPrice": {
                    "type": "number"
                },
                "productCode": {
                    "type": "string"
                },
                "productName": {
                    "type": "string"
                },
                "standadCost": {
                    "type": "number"
                }
            }
        },
        "product.ProductList": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/product.Product"
                    }
                },
                "totalRecords": {
                    "type": "integer"
                }
            }
        },
        "product.ProductTop": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "product_name": {
                    "type": "string"
                },
                "vendidos": {
                    "type": "number"
                }
            }
        },
        "product.getAddProductRequest": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "listPrice": {
                    "type": "string"
                },
                "productCode": {
                    "type": "string"
                },
                "productName": {
                    "type": "string"
                },
                "standardCost": {
                    "type": "string"
                }
            }
        },
        "product.getProductsRequest": {
            "type": "object",
            "properties": {
                "limit": {
                    "type": "integer"
                },
                "offset": {
                    "type": "integer"
                }
            }
        },
        "product.updateProductRequest": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "listPrice": {
                    "type": "number"
                },
                "productCode": {
                    "type": "string"
                },
                "productName": {
                    "type": "string"
                },
                "standardCost": {
                    "type": "number"
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
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Engrenelog Goolivery Provider APi",
	Description: "APIs desarrolladas en Golang para el Sistema Goolivery Provider",
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
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
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
	swag.Register("swagger", &s{})
}