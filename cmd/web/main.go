package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/note", showNote)
	mux.HandleFunc("/note/create", createNote)

	log.Println("Web server is running")
	err := http.ListenAndServe(":9000", mux)
	log.Fatal(err)
}
