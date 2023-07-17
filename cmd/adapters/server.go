package adapters

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	"golang.org/x/exp/slog"

	"github.com/ak2ie/golang_tutorial/cmd/hello"
	"github.com/ak2ie/golang_tutorial/models"
)

// ServerInterfaceを実装

type Server struct{}

func (s *Server) Hello(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("pgx", "host=postgres port=5432 user=postgres dbname=db password=password sslmode=disable")
	if nil != err {
		slog.Error("db connection failed")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	defer cancel()
	err = db.PingContext(ctx)
	if nil != err {
		slog.Error("db cannot open." + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	books, err := models.Books().All(ctx, db)
	if nil != err {
		slog.Error("db select error")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	for i := range books {
		fmt.Printf("%s\n", books[i].Name)
	}

	var (
		id   int
		name string
	)
	row := db.QueryRowContext(ctx, `SELECT id, name FROM book WHERE id = $1`, 1)
	err = row.Scan(&id, &name)
	if nil != err {
		slog.Error("db select error")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "Hello "+name)
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
