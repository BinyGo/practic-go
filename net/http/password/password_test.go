package password

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPassWordHashingHandler(t *testing.T) {

	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		PassWordHashingHandler(w, r)
		w.WriteHeader(200)
		w.Write(nil)
	}

	req := httptest.NewRequest("GET", "http://biny.com", nil)
	w := httptest.NewRecorder()
	handler(w, req)
}
