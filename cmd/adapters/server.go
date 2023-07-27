package adapters

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/jackc/pgx/v4/stdlib"
	"golang.org/x/exp/slog"

	"github.com/ak2ie/golang_tutorial/cmd/hello"
	"github.com/ak2ie/golang_tutorial/generated"
)

// ServerInterfaceを実装
type Server struct{}

// Bookテーブル
// type Book struct {
// 	ID    int
// 	Name  string
// 	Price int
// }

func (s *Server) Hello(w http.ResponseWriter, r *http.Request) {
	dsn := "host=postgres port=5432 user=postgres dbname=db password=password sslmode=disable"
	// db, err := sql.Open("pgx", "host=postgres port=5432 user=postgres dbname=db password=password sslmode=disable")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if nil != err {
		slog.Error("db connection failed")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)
	// defer cancel()
	// err = db.PingContext(ctx)
	// if nil != err {
	// 	slog.Error("db cannot open." + err.Error())
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }

	err = db.AutoMigrate(&generated.Book{})
	if nil != err {
		slog.Error("db migration error! " + err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var book generated.Book
	db.First(&book, 1)

	fmt.Printf("%s %d円\n", book.Name.String, book.Price.Int64)

	fmt.Fprint(w, "Hello!")
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
