package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/yodalis/golang/10-multithreading/desafio/internal/infra/api"
	"github.com/yodalis/golang/10-multithreading/desafio/internal/infra/webservers/handlers"
)

func main() {
	apiViaCep := api.NewViaCEP()
	brasilCep := api.NewBrasilCep()
	addressHandler := handlers.NewAddressHandler(apiViaCep, brasilCep)

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/address/{cep}", addressHandler.GetAddressByCep)

	http.ListenAndServe(":8080", r)
}
