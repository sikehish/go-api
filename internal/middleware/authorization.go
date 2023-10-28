package middleware

import (
	"errors"
	"net/http"

	"github.com/sikehish/go-api/api"
	"github.com/sikehish/go-api/internal/tools"
	log "github.com/sirupsen/logrus"
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
		var err error
		database, err=tools.NewDatabase()

		if err!=nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails=(*database).GetUserLoginDetails(username)

		if (loginDetails==nil || (token!=(*loginDetails).AuthToken)){
			log.Error(unAuthorizedError)
			api.RequestErrorHandler(w, unAuthorizedError)
			return
		}

		next.ServeHTTP(w,r)

	})
}