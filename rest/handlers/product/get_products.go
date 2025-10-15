package product

import (
	"ecommerce/util"
	"net/http"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.productRepo.List()
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	util.SendData(w, products, http.StatusOK)
}
