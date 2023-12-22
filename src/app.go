package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func logPanic(err error) {
	if err != nil {
		log.Panicf("Error Occured :%v", err.Error())
	}
}

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (app *App) Initialize() {
	var err error
	connectionString := fmt.Sprintf("%v:%v@tcp(127.0.0.1:3306)/%v", DbUser, DbPass, DbName)
	app.DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	logPanic(err)
	app.DB.AutoMigrate(&Product{})
	app.Router = mux.NewRouter().StrictSlash(true)

	// return nil

}

func (app *App) Run(address string) error {
	app.Router.HandleFunc("/", app.homePage).Methods("GET")
	app.Router.HandleFunc("/products", app.getProducts).Methods("GET")
	app.Router.HandleFunc("/product/{id}", app.getProduct).Methods("GET")
	app.Router.HandleFunc("/products", app.createProduct).Methods("POST")
	app.Router.HandleFunc("/product/{id}", app.updateProduct).Methods("PUT")
	app.Router.HandleFunc("/product/{id}", app.deleteProduct).Methods("DELETE")
	return http.ListenAndServe(address, app.Router)
}

func (app *App) homePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}
func (app *App) getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	getProductsFromDb(app, w, r)
	w.WriteHeader(http.StatusOK)
}
func (app *App) getProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	getProductFromDb(app, w, r)
}
func (app *App) createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	addProductToDb(app, w, r)

}

func (app *App) updateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	updateProductInDb(app, w, r)
}
func (app *App) deleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	deleteProductInDb(app, w, r)
}
