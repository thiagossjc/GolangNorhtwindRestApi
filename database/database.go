package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func InitDB() *sql.DB {
	connectionString := "root:admin@tcp(localhost:3306)/northwind"
	//fmt.Println(connectionString)
	databaseConnection, err := sql.Open("mysql", connectionString)
	if err != nil {
		panic(err.Error()) //manejo de errores

	}

	return databaseConnection
}
