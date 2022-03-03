package database

import (
	"database/sql"
	"fmt"

	"github.com/GoGooliveryProviderAPI/helper"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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

//Connect DB to Posgresql
func InitDbPgGorm() (*gorm.DB, error) {

	dsn := "host=localhost user=postgres password=1234567 dbname=goolivery_provider port=5432 sslmode=disable TimeZone=Europe/Madrid"
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	helper.Catch(err)
	return DB, err
}
