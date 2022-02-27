package main

import (
	"fmt"
	"log"
	"net/http"
)

type likesStuff struct {
	Name string `json:"name"`

	Likes string `json:"likes"`

	Dislikes string `json:"dislikes"`

	Skills string `json:"skills"`
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/json" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

}

func main() {

	fileServer := http.FileServer(http.Dir("./"))
	http.Handle("/", fileServer)
	http.HandleFunc("/json", jsonHandler)

	fmt.Printf("Starting server at port 1338\n")
	if err := http.ListenAndServe(":1338", nil); err != nil {
		log.Fatal(err)
	}

}
