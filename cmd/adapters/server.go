package adapters

import (
	"fmt"
	"net/http"
)

// ServerInterfaceを実装

type Server struct{}

func (s *Server) Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func NewServer() *Server {
	return &Server{}
}
