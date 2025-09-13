package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser database.User


	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newUser)
	if err != nil {
		http.Error(w, "plz give me a valid json", http.StatusBadRequest)
		return
	}
	newUser = newUser.Store()

	util.SendData(w, newUser, http.StatusCreated)

}
