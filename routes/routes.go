package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func pong(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}

func Routes(router *mux.Router) {
	router.HandleFunc("/ping", pong).Methods("GET")
}
