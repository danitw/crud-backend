package main

import (
	"crud-backend/database"
	"crud-backend/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

const PORT string = "8000"

func main() {
	db := database.Connection()

	database.Migrate(db)

	router := mux.NewRouter().StrictSlash(true)

	router.Use(mux.CORSMethodMiddleware(router))

	routes.Routes(router)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"*"},
	})

	log.Println("initializing server in port", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, c.Handler(router)))
}
