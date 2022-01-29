package main

import (
	"net/http"

	"github.com/GolangNorhtwindRestApi/customer"
	"github.com/GolangNorhtwindRestApi/database"
	"github.com/GolangNorhtwindRestApi/employee"
	"github.com/GolangNorhtwindRestApi/order"
	"github.com/GolangNorhtwindRestApi/product"

	_ "github.com/GolangNorhtwindRestApi/docs"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Engrenelog Goolivery Provider APi
// @version 1.0
// @description APIs desarrolladas en Golang para el Sistema Goolivery Provider
// @contact.name API Support (Thiago Mota)
// @contact.url http://gooliveryprovider
// @contact.email thiago@engrenelog.com

func main() {
	databaseConnection := database.InitDB()
	defer databaseConnection.Close()

	var (
		productRepository  = product.NewRepository(databaseConnection)
		employeeRepository = employee.NewRepository(databaseConnection)
		customerRepository = customer.NewRepository(databaseConnection)
		orderRepository    = order.NewRepository(databaseConnection)
	)
	var (
		productService  product.Service
		employeeService employee.Service
		customerService customer.Service
		orderService    order.Service
	)
	productService = product.NewService(productRepository)
	employeeService = employee.NewService(employeeRepository)
	customerService = customer.NewService(customerRepository)
	orderService = order.NewService(orderRepository)

	r := chi.NewRouter()
	r.Mount("/products", product.MakeHttpHandler(productService))
	r.Mount("/employees", employee.MakeHttpHandler(employeeService))
	r.Mount("/customers", customer.MakeHttpHandler(customerService))
	r.Mount("/orders", order.MakeHttpHandler(orderService))

	r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL("../swagger/doc.json")))
	http.ListenAndServe(":3000", r)
}
