package product

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")

	pId, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var newProduct database.Product

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Plz Give me valid json", 400)
		return
	}

	newProduct.ID = pId

	updatedProduct := database.Update(newProduct)

	if updatedProduct == nil {
		util.SendError(w, http.StatusNotFound, "Product not found")
		return
	}

	util.SendData(w, "successfully updated product", http.StatusOK)
}
