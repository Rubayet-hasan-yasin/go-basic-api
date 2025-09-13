package review

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
)

func (h *Handler) GetReview(w http.ResponseWriter, r *http.Request) {

	util.SendData(w, database.List(), http.StatusOK)
}
