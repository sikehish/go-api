package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sikehish/go-api/internal/handlers"
	log "github.com/sirupsen/logrus"
)


func main(){
	log.SetReportCaller(true) //SetReportCaller(true) is used to enable reporting of the file and line number where a log message is generated.
	var r *chi.Mux=chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Starting GO API service...")

	err:=http.ListenAndServe("localhost:8080", r)
	if err!=nil {
		log.Error(err)
	}
	
}