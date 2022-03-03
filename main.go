package main

import (
	"fmt"
	"net/http"

	"github.com/GolangNorhtwindRestApi/customer"
	"github.com/GolangNorhtwindRestApi/database"
	"github.com/GolangNorhtwindRestApi/employee"
	"github.com/GolangNorhtwindRestApi/helper"
	"github.com/GolangNorhtwindRestApi/order"
	"github.com/GolangNorhtwindRestApi/product"
	"github.com/GolangNorhtwindRestApi/status"

	_ "github.com/GolangNorhtwindRestApi/docs"
	"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Engrenelog Goolivery Provider APi
// @version 1.0
// @description APIs desarrolladas en Golang para el Sistema Goolivery Provider
// @contact.name API Support (Thiago Mota)
// @contact.url http://engrenelog.com/es/
// @contact.email thiago@engrenelog.com
func main() {

	databasePGGormConnection, err := database.InitDbPgGorm()
	if err == nil {
		fmt.Println("Conexão con el Banco de Dados PostgreSql ok!")
	} else {
		fmt.Println(err)
	}
	defer databasePGGormConnection.Statement.ReflectValue.Close()

	databasePGConnection, err := database.InitDBPG()
	if err == nil {
		fmt.Println("Conexão con el Banco de Dados PostgreSql ok!")
	} else {
		fmt.Println(err)
	}
	defer databasePGConnection.Close()

	databaseConnection, err := database.InitDB()
	if err == nil {
		fmt.Println("Conexão con el Banco de Dados MariaDB ok!")
	} else {
		fmt.Println(err)
	}
	defer databaseConnection.Close()

	var (
		productRepository  = product.NewRepository(databaseConnection)
		employeeRepository = employee.NewRepository(databaseConnection)
		customerRepository = customer.NewRepository(databaseConnection)
		orderRepository    = order.NewRepository(databaseConnection)
		statusRepository   = status.NewRepository(databasePGGormConnection)
	)
	var (
		productService  product.Service
		employeeService employee.Service
		customerService customer.Service
		orderService    order.Service
		statusService   status.Service
	)
	productService = product.NewService(productRepository)
	employeeService = employee.NewService(employeeRepository)
	customerService = customer.NewService(customerRepository)
	orderService = order.NewService(orderRepository)
	statusService = status.NewService(statusRepository)

	r := chi.NewRouter()
	r.Use(helper.GetCors().Handler) //Invocanco CORSxº

	r.Mount("/products", product.MakeHttpHandler(productService))
	r.Mount("/employees", employee.MakeHttpHandler(employeeService))
	r.Mount("/customers", customer.MakeHttpHandler(customerService))
	r.Mount("/orders", order.MakeHttpHandler(orderService))
	r.Mount("status", status.MakeHttpHandler(statusService))

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("../swagger/doc.json"),
	))
	http.ListenAndServe(":3000", r)
}
