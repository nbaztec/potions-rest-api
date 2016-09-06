package main

import (
	"github.com/nishb/test-potions/routes"
	"net/http"
)

func main() {
	http.ListenAndServe(":9100", routes.NewRouter())
}
