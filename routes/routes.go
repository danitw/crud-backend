package routes

import (
	"crud-backend/controllers"
	"crud-backend/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

func pong(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}

func Routes(router *mux.Router) {
	router.HandleFunc("/register", controllers.Register).Methods("POST")
	router.HandleFunc("/login", controllers.Login).Methods("POST")

	router.Handle("/post", middlewares.JWT(http.HandlerFunc(controllers.CreatePost))).Methods("POST")
	router.HandleFunc("/post/{post_id:[0-9]+}", controllers.ReadPost).Methods("GET")
	router.HandleFunc("/post", controllers.ListPost).Methods("GET")
	router.Handle("/post/{post_id:[0-9]+}", middlewares.JWT(http.HandlerFunc(controllers.DeletePost))).Methods("DELETE")
	router.Handle("/post/{post_id:[0-9]+}", middlewares.JWT(http.HandlerFunc(controllers.UpdatePost))).Methods("PUT")
}
