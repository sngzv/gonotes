package main

import (
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Gonotes"))
}

func showNote(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Show note..."))
}

func createNote(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Form for create a new note..."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("note", showNote)
	mux.HandleFunc("note/create", createNote)

	log.Println("Web server is running")
	err := http.ListenAndServe(":9000", mux)
	log.Fatal(err)
}
