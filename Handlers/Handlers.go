package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func JsonHandler(w http.ResponseWriter, r *http.Request) {
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

func BoopHandler(w http.ResponseWriter, r *http.Request) {
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
