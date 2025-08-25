package handlers

import (
	"e-commerce/database"
	"e-commerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	// allowCors(w)

	// Options(w, r)

	// POST(w, r)

	var newProduct database.Product

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "plz give me a valid json", http.StatusBadRequest)
		return
	}

	newProduct.ID = len(database.Products) + 1

	database.Products = append(database.Products, newProduct)

	util.SendData(w, newProduct, http.StatusCreated)
}