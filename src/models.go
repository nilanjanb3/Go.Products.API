package main

import (
	"encoding/json"
	"net/http"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Id       string
	Name     string
	Quantity int
	Price    float64
}

func addProductToDb(app *App, w http.ResponseWriter, r *http.Request) {
	var product Product
	_ = json.NewDecoder(r.Body).Decode(&product)
	app.DB.Create(&product)
}
