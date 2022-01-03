package main

import (
	"GolangNorhtwindRestApi/database"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var databaseConnection *sql.DB

type Product struct {
	ID           int    `json:"id"`
	Product_Code string `json:"product_code"`
	Description  string `json:"description"`
}

func catch(err error) {
	if err != nil {
		panic(err)
	}

}

func main() {
	databaseConnection := database.InitDB()
	defer databaseConnection.Close()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/products", AllProductos)
	r.Post("/products", CreateProducto)
	r.Put("/products/{id}", UpdateProducto)
	r.Delete("/products/{id}", DeleteProducto)
	http.ListenAndServe(":3000", r)
}

func DeleteProducto(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	query, err := databaseConnection.Prepare("delete from products where id=?")
	catch(err)
	_, er := query.Exec(id)
	catch(er)
	query.Close()
	respondiwithJSON(w, http.StatusOK, map[string]string{"message": "successfully deleted"})
}

func UpdateProducto(w http.ResponseWriter, r *http.Request) {
	var producto Product
	id := chi.URLParam(r, "id")
	json.NewDecoder(r.Body).Decode(&producto)
	query, err := databaseConnection.Prepare("Update products SET product_cod=?,description=? where id=?")
	catch(err)
	_, er := query.Exec(producto.Product_Code, producto.Description, id)
	catch(er)
	defer query.Close()
	respondiwithJSON(w, http.StatusCreated, map[string]string{"message": "update created"})
}

func CreateProducto(w http.ResponseWriter, r *http.Request) {
	var producto Product
	json.NewDecoder(r.Body).Decode(&producto)
	query, err := databaseConnection.Prepare("Insert products SET product_cod=?,description=?")
	catch(err)
	_, er := query.Exec(producto.Product_Code, producto.Description)
	catch(er)
	defer query.Close()
	respondiwithJSON(w, http.StatusCreated, map[string]string{"message": "successfully created"})
}

func AllProductos(w http.ResponseWriter, r *http.Request) {
	const sql = `SELECT id, product_code, COALESCE(description,'')
				from products`
	results, err := databaseConnection.Query(sql)
	catch(err)
	var products []*Product //slice

	for results.Next() {
		product := &Product{} //Puntero
		err = results.Scan(&product.ID, &product.Product_Code, &product.Description)

		catch(err)
		products = append(products, product)
	}
	respondiwithJSON(w, http.StatusOK, products)
}

func respondiwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	fmt.Println(payload)
	w.Header().Set("Context-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
