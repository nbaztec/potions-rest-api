package routes

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"bytes"
)

func TestGetRoute(t *testing.T) {
	r, _ := http.NewRequest("GET", "/get", nil)
	w := httptest.NewRecorder()
	GetRoute(w, r)
	if w.Code != http.StatusBadRequest {
		t.Error("did not receive HTTP 400")
	}

	r, _ = http.NewRequest("GET", "/get?id=foo", nil)
	w = httptest.NewRecorder()
	GetRoute(w, r)
	if w.Code != http.StatusBadRequest {
		t.Error("did not receive HTTP 400")
	}

	r, _ = http.NewRequest("GET", "/get?id=0", nil)
	w = httptest.NewRecorder()
	GetRoute(w, r)
	if w.Code != http.StatusNotFound {
		t.Error("did not receive HTTP 404")
	}

	resetItems()
	r, _ = http.NewRequest("GET", "/get?id=2", nil)
	w = httptest.NewRecorder()
	GetRoute(w, r)
	if w.Code != http.StatusOK {
		t.Error("did not receive HTTP 200")
	}
}

func TestMixRoute(t *testing.T) {
	r, _ := http.NewRequest("POST", "/mix", bytes.NewReader([]byte(`{}`)))
	w := httptest.NewRecorder()
	MixPotionRoute(w, r)
	if w.Code != http.StatusBadRequest {
		t.Error("did not receive HTTP 400")
	}

	r, _ = http.NewRequest("POST", "/mix", bytes.NewReader([]byte(`{ids: []}`)))
	w = httptest.NewRecorder()
	MixPotionRoute(w, r)
	if w.Code != http.StatusBadRequest {
		t.Error("did not receive HTTP 400")
	}

	resetItems()

	r, _ = http.NewRequest("POST", "/mix", bytes.NewReader([]byte(`{ids: [1]}`)))
	w = httptest.NewRecorder()
	MixPotionRoute(w, r)
	if w.Code != http.StatusBadRequest {
		t.Error("did not receive HTTP 400")
	}

	r, _ = http.NewRequest("POST", "/mix", bytes.NewReader([]byte(`{"ids": [0,1,2]}`)))
	w = httptest.NewRecorder()
	MixPotionRoute(w, r)
	if w.Code != http.StatusNotFound {
		t.Error("did not receive HTTP 404")
	}

	r, _ = http.NewRequest("POST", "/mix", bytes.NewReader([]byte(`{"ids": [1,2]}`)))
	w = httptest.NewRecorder()
	MixPotionRoute(w, r)
	if w.Code != http.StatusOK {
		t.Error("did not receive HTTP 200")
	}

	r, _ = http.NewRequest("POST", "/mix", bytes.NewReader([]byte(`{"ids": ["1","2"]}`)))
	w = httptest.NewRecorder()
	MixPotionRoute(w, r)
	if w.Code != http.StatusOK {
		t.Error("did not receive HTTP 200")
	}

	r, _ = http.NewRequest("POST", "/mix", bytes.NewReader([]byte(`{"ids": [1, "a"]}`)))
	w = httptest.NewRecorder()
	MixPotionRoute(w, r)
	if w.Code != http.StatusBadRequest {
		t.Error("did not receive HTTP 400")
	}

	r, _ = http.NewRequest("POST", "/mix", bytes.NewReader([]byte(`{"ids": [1, false]}`)))
	w = httptest.NewRecorder()
	MixPotionRoute(w, r)
	if w.Code != http.StatusBadRequest {
		t.Error("did not receive HTTP 400")
	}
}