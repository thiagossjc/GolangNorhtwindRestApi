package product

import "dabase/sql"

type Repository interface {
	GetProductoById(productoID int) (*Product, error)
}

type repository struct {
	db *sql.DB
}

func newRepository(databaseConnection *sql.db) Repository
{
	return &repository{db:databaseConnection}
}

func (repo *repository) GetProductoById(productoId int) (*Product, error){

	const sql = `SELECT id	product_cod, product_name,COALESCE(description,''), standard_cost,category from products where id=?`
	row:= repo.db.QueryRow(sql,productoId)
	product:= &Product{}
	err:= row.Scan(&product.Id,&product.ProductCode,&product.ProductName,
	&product.Description,&product.ListPrice,&product.Category)

	if err != nil{
		panic(err)
	}
	return product,err
}