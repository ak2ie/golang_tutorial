package main

import (
	"log"
	"net/http"

	"github.com/ak2ie/golang_tutorial/cmd/adapters"
	"github.com/ak2ie/golang_tutorial/cmd/hello"
	middleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/go-chi/chi/v5"
	"golang.org/x/exp/slog"
)

func main() {
	swagger, _ := hello.GetSwagger()
	swagger.Servers = nil
	r := chi.NewRouter()
	r.Use(middleware.OapiRequestValidator(swagger))
	h := adapters.NewServer()
	hello.HandlerFromMux(h, r)
	log.Fatal(http.ListenAndServe(":8080", MiddlewareLogging(r)))
}

// ロギング
func MiddlewareLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("start", "URL", r.URL)
		lrw := NewLoggingResposeWriter(w)

		next.ServeHTTP(lrw, r)

		statusCode := lrw.statusCode
		slog.Info("end", "URL", r.URL, "StatusCode", http.StatusText(statusCode))
	})
}

// ResponseWriterを満たす
type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLoggingResposeWriter(w http.ResponseWriter) *loggingResponseWriter {
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code // ステータスコードを保存
	lrw.ResponseWriter.WriteHeader(code)
}
