package database

import (
	"database/sql"

	"github.com/GolangNorhtwindRestApi/helper"
	_ "github.com/go-sql-driver/mysql"
	//_ "github.com/go-sql-driver/mariadb"
)

func InitDB() (*sql.DB, error) {
	connectionString := "root:admin@tcp(localhost:3306)/northwind"
	//fmt.Println(connectionString)
	databaseConnection, err := sql.Open("mysql", connectionString)
	helper.Catch(err)

	return databaseConnection, err
}
