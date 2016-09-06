package main

import (
	"github.com/nishb/potions-rest-api/routes"
	"net/http"
)

func main() {
	http.ListenAndServe(":9100", routes.NewRouter())
}
