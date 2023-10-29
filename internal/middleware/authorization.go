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

		//NOTE:
		// In Go, http.ResponseWriter is not typically taken as a pointer because it is an interface. Go interfaces are already reference types, and you don't need a pointer to modify the value that implements an interface.  You can pass an interface by value (not as a pointer) and still make modifications to the underlying object.you use it to write the response (e.g., writing headers and response content). The implementations of the interface take care of modifying the response based on the methods you call.

	
	// In this code, the Authorization middleware function is designed to return an http.Handler that performs authorization checks before allowing the actual request handling to proceed. Let's break down how this works:

    // http.HandlerFunc is an adapter function. It converts a regular Go function with the signature func(http.ResponseWriter, *http.Request) into an http.Handler, which is necessary for middleware to fit into the Go net/http middleware chain.

    // The Authorization middleware function is defined to return an http.Handler. It takes another http.Handler (next) as an argument, which is typically the next handler in the chain, and wraps it with authorization logic.

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

			//Note: n Go, both -> and . are represented by ""." .The compiler knows the types, and can dereference if necessary.
			//loginDetails.AuthToken == (*loginDetails).AuthToken
		if (loginDetails==nil || (token!=loginDetails.AuthToken)){
			log.Error(unAuthorizedError)
			api.RequestErrorHandler(w, unAuthorizedError)
			return
		}

		next.ServeHTTP(w,r)
		// The next.ServeHTTP(w, r) call is used to pass the control of the HTTP request to the next handler in the middleware chain. It allows the request to continue processing by the subsequent handlers in the chain.
		// When you call next.ServeHTTP(w, r) within the Authorization middleware, you're effectively passing the http.ResponseWriter and *http.Request objects to the next handler in the middleware chain. This means that after the Authorization middleware performs its checks and logic, it allows the subsequent handler (e.g., your application's request handler) to process the request.

	})
}