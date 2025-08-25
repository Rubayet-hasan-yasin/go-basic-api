package handlers

import (
	"e-commerce/database"
	"e-commerce/util"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	// allowCors(w)

	// GET(w, r)

	util.SendData(w, database.Products, http.StatusOK)
}
