package adapters

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ak2ie/golang_tutorial/cmd/hello"
)

// ServerInterfaceを実装

type Server struct{}

func (s *Server) Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func (s *Server) PostHello(w http.ResponseWriter, r *http.Request) {
	var c hello.Sample

	if err := JsonDecode[hello.Sample](r, &c); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(*c.Id))
}

func NewServer() *Server {
	return &Server{}
}

func JsonDecode[T any](r *http.Request, inputData *T) error {
	defer r.Body.Close()

	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields() // 指定外のフィールド禁止
	if err := d.Decode(inputData); err != nil {
		return err
	}
	return nil
}
