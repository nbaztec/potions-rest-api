package routes

import (
	"bytes"
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Test struct {
	path    string
	methods []string
}

var testRoutes = []Test{
	{"/", []string{"GET", "POST"}},
	{"/list", []string{"GET"}},
	{"/item", []string{"GET"}},
	{"/mix", []string{"POST"}},
}

func TestNewRouterRoutes(t *testing.T) {
	router := NewRouter()
	for _, route := range testRoutes {
		for _, m := range route.methods {
			w := httptest.NewRecorder()
			r, _ := http.NewRequest(m, route.path, bytes.NewReader([]byte(`{}`)))
			r.Header.Set("Authorization", base64.StdEncoding.EncodeToString([]byte(authKey)))
			router.ServeHTTP(w, r)
			if w.Code == http.StatusForbidden {
				t.Error("received HTTP 404")
			}
		}
	}
}

func TestRoute404(t *testing.T) {
	router := NewRouter()
	r, _ := http.NewRequest("GET", "/path/not/exists", nil)
	r.Header.Set("Authorization", base64.StdEncoding.EncodeToString([]byte(authKey)))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)
	if w.Code != http.StatusNotFound {
		t.Error("did not receive HTTP 404")
	}
}