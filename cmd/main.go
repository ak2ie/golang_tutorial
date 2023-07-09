package main

import (
	"net/http"

	"github.com/ak2ie/golang_tutorial/cmd/adapters"
	"github.com/ak2ie/golang_tutorial/cmd/hello"
)

func main() {
	http.ListenAndServe(":8080", hello.Handler(&adapters.Server{}))
}
