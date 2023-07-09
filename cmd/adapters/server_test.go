package adapters

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/hello", nil)
	w := httptest.NewRecorder()

	s := NewServer()
	s.Hello(w, r)

	if w.Body.String() != "hello world!" {
		t.Errorf("result should be \"hello world!\", but %s", w.Body.String())
	}
}
