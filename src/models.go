package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	// gorm.Model
	Id       string
	Name     string
	Quantity int
	Price    float64
}

func getProductsFromDb(app *App, w http.ResponseWriter, r *http.Request) {
	var products []Product
	app.DB.Find(&products)
	json.NewEncoder(w).Encode(products)
}

func getProductFromDb(app *App, w http.ResponseWriter, r *http.Request) {
	var product Product
	app.DB.First(&product, mux.Vars(r)["id"])
	if product.Id == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Product Not Found")
	} else {
		json.NewEncoder(w).Encode(product)
	}

}
func addProductToDb(app *App, w http.ResponseWriter, r *http.Request) {
	var product Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	app.DB.Create(&product)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Product Created")
}

func updateProductInDb(app *App, w http.ResponseWriter, r *http.Request) {
	var oldProduct Product
	var newProduct Product
	app.DB.First(&oldProduct, mux.Vars(r)["id"])
	_ = json.NewDecoder(r.Body).Decode(&newProduct)
	if oldProduct.Id == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Product Not Found")
	} else if newProduct.Id != mux.Vars(r)["id"] {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "Forbidden")

	} else {

		app.DB.Save(&newProduct)
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(w, "Product Updated")
	}

}

func deleteProductInDb(app *App, w http.ResponseWriter, r *http.Request) {
	var product Product
	app.DB.First(&product, mux.Vars(r)["id"])
	if product.Id == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Product Not Found")
	} else {
		w.WriteHeader(http.StatusAccepted)
		app.DB.Delete(&product)
		fmt.Fprintf(w, "Product Deleted")
	}

}
