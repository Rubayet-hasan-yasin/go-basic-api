package user

import (
	"ecommerce/config"
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)


type ReqLogin struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request){
	var reqLogin ReqLogin
	decoder := json.NewDecoder((r.Body))
	err := decoder.Decode(&reqLogin)
	if err != nil{
		fmt.Println(err)
		util.SendError(w, http.StatusBadRequest, "Invalid Request Data")
		return
	}

	usr := database.Find(reqLogin.Email, reqLogin.Password)
	if usr == nil{
		util.SendError(w, http.StatusBadRequest, "Invalid Credentials")
		return
	}

	cnf := config.GetConfig()
	accessToken, err := util.CreateJwt(cnf.JwtSecretKey, util.Payload{
		Sub: usr.ID,
		FirstName: usr.FirstName,
		LastName: usr.LastName,
		Email: usr.Email,
		IsShopOwner: usr.IsShopOwner,
	})
	if err != nil{
		util.SendError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	util.SendData(w, accessToken, http.StatusCreated)
}