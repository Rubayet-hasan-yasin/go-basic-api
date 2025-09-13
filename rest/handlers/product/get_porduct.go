package product

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)


func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")

	pId, err := strconv.Atoi(productId)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product := database.GetByID(pId)
	if product == nil {
		util.SendError(w, http.StatusNotFound ,"Product not found")
		return
	}

	util.SendData(w, product, http.StatusOK)
}
