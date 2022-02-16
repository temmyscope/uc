package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type Clip struct {
	id        int    `json:"id"`
	userId    int    `json:"userId"`
	text      string `json:"text"`
	createdAt string `json:"createdAt"`
}

type ClipBoard []Clip

//route for fetching content to clipBoard
func fetchClip(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ClipBoard Content:")
}

//route for adding content to clipBoard
func addClip(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Home")
}

// route for adding content to clipBoard
// (Run in background) for syncing offline clipboard with server
func addClips(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Home")
}

// route for updating offline clipboard from server
func fetchClips(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Home")
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Home")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/sync/clips", fetchClips).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	godotenv.Load()
	handleRequests()
}
