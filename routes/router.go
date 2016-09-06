package routes

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() http.Handler {
	resetItems()

	r := mux.NewRouter()
	r.Methods("GET", "POST").Path("/").HandlerFunc(HomeRoute)
	r.Methods("GET").Path("/list").HandlerFunc(ListRoute)
	r.Methods("GET").Path("/item").HandlerFunc(GetRoute)
	r.Methods("POST").Path("/mix").HandlerFunc(MixPotionRoute)
	r.PathPrefix("/").HandlerFunc(Route404)

	return checkAuth(handlers.CORS()(r))
}
