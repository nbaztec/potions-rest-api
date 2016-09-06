package routes

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func sendReponse(w http.ResponseWriter, status int, r map[string]interface{}) {
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(r); err != nil {
		log.Println("json:", err)
	}
}

func parseJSONBody(r io.Reader) (map[string]interface{}, error) {
	params := make(map[string]interface{})
	if err := json.NewDecoder(r).Decode(&params); err != nil {
		return nil, err
	}
	return params, nil
}
