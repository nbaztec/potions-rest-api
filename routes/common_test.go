package routes

import (
	"encoding/json"
	"net/http/httptest"
	"testing"
)

func TestSendResponse(t *testing.T) {
	w := httptest.NewRecorder()
	sendReponse(w, 200, map[string]interface{}{"status": "Ok"})
	var v map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &v)
	if err != nil {
		t.Error("error decoding json")
	}
	if v["status"] != "Ok" {
		t.Error("invalid json decode result")
	}
}
