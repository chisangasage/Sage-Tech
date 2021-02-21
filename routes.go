package main

import (
	"Sage-Tech/web"
	"net/http"

	"github.com/gorilla/mux"
)

func router() *mux.Router {
	r := mux.NewRouter()

	fs := http.FileServer(http.Dir("./assets/"))
	statichandler := http.StripPrefix("/assets/", fs)
	r.PathPrefix("/assets/").Handler(statichandler)
	r.HandleFunc("/", web.IndexHandler).Methods("GET")
	return r
}
