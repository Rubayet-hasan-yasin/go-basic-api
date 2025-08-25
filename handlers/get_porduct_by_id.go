package handlers

import (
	"e-commerce/database"
	"e-commerce/util"
	"net/http"
	"strconv"
)

func GetProductById(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")

	pId, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	for _, product := range database.Products {
		if product.ID == pId {
			util.SendData(w, product, http.StatusOK)
			return
		}
	}

	http.Error(w, "Product not found", http.StatusNotFound)
	// util.SendData(w, "Product not found", http.StatusNotFound)
}
