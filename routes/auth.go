package routes

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
)

const authKey = "mixOfRandomPotions"

func writeUnauthorized(w http.ResponseWriter) {
	w.WriteHeader(http.StatusForbidden)
	if err := json.NewEncoder(w).Encode(map[string]string{
		"error": "Unauthorized access",
	}); err != nil {
		log.Println("json:", err)
	}
}

func checkAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authStrBase64 := r.Header.Get("Authorization")
		if authStrBase64 == "" {
			writeUnauthorized(w)
			return
		}

		authStr, err := base64.StdEncoding.DecodeString(authStrBase64)
		if err != nil {
			writeUnauthorized(w)
			return
		}

		if authKey != string(authStr) {
			writeUnauthorized(w)
			return
		}

		next.ServeHTTP(w, r)
	})
}
