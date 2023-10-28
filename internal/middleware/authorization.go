package middleware

import (
	"errors"
	"net/http"

	"github.com/sikehish/go-api/api"
)

var unAuthorizedError= errors.New("Invalid username or token.")

//Authorization middleware
func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		//Get username and token from header
		var username string=r.URL.Query().Get("username")
		var token=r.Header.Get("Authorization")

		//Validate username and token
		if username=="" || token=="" {
			log.Error(unAuthorizedError)
			api.RequestErrorHandler(w, unAuthorizedError)
		}

		var database *tools.DatabaseInterface
		database, err=tools.NewDatabase()

		if err!=nil {
			api.InternalErrorHandler(e)
			return
		}

		var loginDetails *tools.loginDetails
		loginDetails, err=(*database).GetLoginDetails(username)

		if (loginDetails==nil || (token!=(*loginDetails).AuthToken)){
			log.Error(unAuthorizedError)
			api.RequestErrorHandler(w, unAuthorizedError)
			return
		}

		next.ServeHTTP(w,r)

	})
}