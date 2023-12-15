package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func HandleError(err error) {
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
	HandleError(err)
	app.DB.AutoMigrate(&Product{})
	app.Router = mux.NewRouter().StrictSlash(true)

	// return nil

}

func (app *App) Run(address string) error {
	app.Router.HandleFunc("/", homePage).Methods("GET")
	app.Router.HandleFunc("/products", app.createProduct).Methods("POST")
	return http.ListenAndServe(address, app.Router)
}

func (app *App) createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	addProductToDb(app, w, r)
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Product Created")

}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}
