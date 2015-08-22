package main

import (
	"log"
	"net/http"
)

func main() {

	var repo = &InMemoryRepository{}
	repo.Init()
	router := NewRouter(repo)
	log.Fatal(http.ListenAndServe(":8080", router))
}