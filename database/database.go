package database

import (
	"database/sql"
	"fmt"

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

func InitDBPG() (*sql.DB, error) {

	const PostgresDriver = "postgres"
	const User = "postgres"
	const Host = "localhost"
	const Port = "5432"
	const Password = "1234567"
	const DbName = "goolivery_provider"

	var DataSourceName = fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname= %s sslmode=disable", Host, Port, User, Password, DbName)

	databaseConnection, err := sql.Open(PostgresDriver, DataSourceName)

	helper.Catch(err)
	return databaseConnection, err
}
