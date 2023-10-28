package api

import (
	"encoding/json"
	"net/http"
)

//Coin Balance params
type CoinBalanceParams struct{
	Username string
}

type CoinBalanceResponse struct{
	Code int //Success code
	Balance int64 //Account balance
}

//Error response
type Error struct{
	Code int //Error code
	Message string //Error message
}

func writeError(w http.ResponseWriter, message string, code int){
	resp:=Error{
		Code: code,
		Message: message,
	}

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}

var(
	RequestErrorHandler=func(w http.ResponseWriter, err error){
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler=func(w http.ResponseWriter){
		writeError(w,"An Unexpected Error Occurred", http.StatusInternalServerError)
	}
)