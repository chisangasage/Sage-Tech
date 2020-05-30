package main

import (
	"log"
	"net/http"
)

func main() {
	r := router()
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
