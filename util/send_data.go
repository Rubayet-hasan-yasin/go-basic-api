package util

import (
	"encoding/json"
	"net/http"
)

func SendData(w http.ResponseWriter, data interface{}, status int) {
	w.WriteHeader(status)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}


func SendError (w http.ResponseWriter, status int, msg string){
	w.WriteHeader(status)
	encoder := json.NewEncoder(w)
	encoder.Encode(msg)
}