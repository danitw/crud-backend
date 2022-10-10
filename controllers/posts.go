package controllers

import (
	"crud-backend/database"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func CreatePost(w http.ResponseWriter, r *http.Request) {
	post := &database.Post{}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responseError{Error: err.Error()})
		return
	}

	err = json.Unmarshal(reqBody, &post)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responseError{Error: err.Error()})
		return
	}

	result := db.Create(post)

	if result.Error != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responseError{Error: result.Error.Error()})
		return
	}

	json.NewEncoder(w).Encode(post)
}

func ReadPost(w http.ResponseWriter, r *http.Request) {
  post := &database.Post{}

	vars := mux.Vars(r)

	db.Find(&post, vars["post_id"])

	json.NewEncoder(w).Encode(post)
}

func ListPost(w http.ResponseWriter, r *http.Request) {
  posts := &[]database.Post{}

	db.Find(&posts)

	json.NewEncoder(w).Encode(posts)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
  post := &database.Post{}
  body := &database.Post{}

	vars := mux.Vars(r)

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(responseError{Error: err.Error()})
		return
	}

	err = json.Unmarshal(reqBody, &body)

	db.First(&post, vars["post_id"]).Updates(body)

	json.NewEncoder(w).Encode(post)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
  post := &database.Post{}

	vars := mux.Vars(r)

	db.Delete(&post, vars["post_id"])

	json.NewEncoder(w).Encode("Deleted with success")
}


