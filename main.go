package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Server", "Go")
	w.Write([]byte("Hello from snippet box"))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("A form for writing snippet"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Displaying snippetbox with id -> %d...", id)
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save a new snippet"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)

	log.Print("Starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
