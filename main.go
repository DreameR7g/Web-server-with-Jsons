package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/json" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	var lotsofStuff string

	stuff, err := os.ReadFile("./jsons/stuff.json")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	_ = json.Unmarshal(stuff, &lotsofStuff)

	w.Write(stuff)

}

func boopHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/boop" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}

	var tonsofBoops string

	boop, err := os.ReadFile("./jsons/boop.json")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	_ = json.Unmarshal(boop, &tonsofBoops)

	w.Write(boop)
}

func main() {

	fileServer := http.FileServer(http.Dir("./"))
	http.Handle("/", fileServer)
	http.HandleFunc("/json", jsonHandler)
	http.HandleFunc("/boop", boopHandler)

	fmt.Printf("Starting server at port 1338\n")
	if err := http.ListenAndServe(":1338", nil); err != nil {
		log.Fatal(err)
	}

}
