package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/sikehish/go-api/internal/middleware"
)

func Handler(r *chi.Mux){
	//Global middleware
	r.Use(chimiddle.StripSlashes) //it strips trailing slashes as it'll lead to 404 error

	r.Route("/account", func(router chi.Router){

		//Middleware for /account route
		router.Use(middleware.Authorization)
		router.Get("/coins", GetCoinBalance)
	})
}

//Difference between chi.Mux vs chi.Router

// chi.Router: chi.Router is an interface in the Chi package that defines the common methods for routing in Chi. It includes methods for defining routes, applying middleware, and other routing features. It serves as the basis for routing in Chi.

// chi.Mux: chi.Mux is the primary and most commonly used router implementation in Chi. It is a concrete type that implements the chi.Router interface. chi.Mux provides a wide range of routing capabilities, including routing, middleware, subrouting, and more. It is suitable for most web application use cases.

// To summarize, chi.Mux is the primary router implementation in Chi, while chi.Router is an interface that defines the common routing methods. You would typically use chi.Mux as your primary router when building web applications with Chi.