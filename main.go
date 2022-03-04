package main

import (
	"fmt"
	"net/http"

	"github.com/GoGooliveryProviderAPI/customer"
	"github.com/GoGooliveryProviderAPI/database"
	"github.com/GoGooliveryProviderAPI/employee"
	"github.com/GoGooliveryProviderAPI/event"
	"github.com/GoGooliveryProviderAPI/helper"
	"github.com/GoGooliveryProviderAPI/order"
	"github.com/GoGooliveryProviderAPI/product"
	"github.com/GoGooliveryProviderAPI/status"
	"github.com/GoGooliveryProviderAPI/user"

	_ "github.com/GoGooliveryProviderAPI/docs"
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
		fmt.Println("Conexão con el Banco de Dados PostgreSql Gorm ok!")
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
		eventsRepository   = event.NewRepository(databasePGGormConnection)
		userRepository     = user.NewRepository(databasePGGormConnection)
	)
	var (
		productService  product.Service
		employeeService employee.Service
		customerService customer.Service
		orderService    order.Service
		statusService   status.Service
		eventsService   event.Service
		userService     user.Service
	)
	productService = product.NewService(productRepository)
	employeeService = employee.NewService(employeeRepository)
	customerService = customer.NewService(customerRepository)
	orderService = order.NewService(orderRepository)
	statusService = status.NewService(statusRepository)
	eventsService = event.NewService(eventsRepository)
	userService = user.NewService(userRepository)

	r := chi.NewRouter()
	r.Use(helper.GetCors().Handler) //Invocanco CORSxº

	r.Mount("/products", product.MakeHttpHandler(productService))
	r.Mount("/employees", employee.MakeHttpHandler(employeeService))
	r.Mount("/customers", customer.MakeHttpHandler(customerService))
	r.Mount("/orders", order.MakeHttpHandler(orderService))
	r.Mount("/status", status.MakeHttpHandler(statusService))
	r.Mount("/events", event.MakeHttpHandler(eventsService))
	r.Mount("/users", user.MakeHttpHandler(userService))

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("../swagger/doc.json"),
	))
	http.ListenAndServe(":3000", r)
}
