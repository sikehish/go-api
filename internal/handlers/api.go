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