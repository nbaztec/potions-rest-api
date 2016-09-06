package routes

import (
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockHTTPHandler struct{}

func (h mockHTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func TestCorrectCredentials(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/", nil)
	r.Header.Set("Authorization", base64.StdEncoding.EncodeToString([]byte(authKey)))

	checkAuth(mockHTTPHandler{}).ServeHTTP(w, r)
	if w.Code != http.StatusOK {
		t.Error("correct credentials do not return HTTP 200")
	}
}

func TestIncorrectCredentials(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/", nil)
	r.Header.Set("Authorization", base64.StdEncoding.EncodeToString([]byte("-")))

	checkAuth(mockHTTPHandler{}).ServeHTTP(w, r)
	if w.Code != http.StatusForbidden {
		t.Error("incorrect credentials do not return HTTP 404")
	}

	w = httptest.NewRecorder()
	r.Header.Set("Authorization", base64.StdEncoding.EncodeToString([]byte("")))

	checkAuth(mockHTTPHandler{}).ServeHTTP(w, r)
	if w.Code != http.StatusForbidden {
		t.Error("empty credentials do not return HTTP 404")
	}

	w = httptest.NewRecorder()
	r.Header.Set("Authorization", "!!")

	checkAuth(mockHTTPHandler{}).ServeHTTP(w, r)
	if w.Code != http.StatusForbidden {
		t.Error("invalid base64 credentials do not return HTTP 404")
	}
}
