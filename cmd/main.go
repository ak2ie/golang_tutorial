package main

import (
	"log"
	"net/http"

	"github.com/ak2ie/golang_tutorial/cmd/adapters"
	"github.com/ak2ie/golang_tutorial/cmd/hello"
	middleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	swagger, _ := hello.GetSwagger()
	swagger.Servers = nil
	r := chi.NewRouter()
	r.Use(middleware.OapiRequestValidator(swagger))
	h := adapters.NewServer()
	hello.HandlerFromMux(h, r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
