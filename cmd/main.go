package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.ListenAndServe(":8080", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	test := "Hello Remote Debug!!!"
	fmt.Fprint(w, test)
}
