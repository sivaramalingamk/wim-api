package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"wim-api/api"
	"wim-api/io"
	"wim-api/services"
)

func main() {
	fmt.Println(" WIM API Serveice is running on Port 8080")
	api.MergeAndInsertTraining()
	io.API_Key = api.ApiKeySetter()
	r := chi.NewRouter()
	services.ProcessInference()
	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	r.Get("/{id},{year}", api.GetOutput)
	r.Route("/ecudata", func(r chi.Router) {
		r.Post("/simple", api.SimpleDataHandler)
		r.Post("/bulk", api.BulkVehicleDataHandler)

	})

	http.ListenAndServe(":8080", r)

}
