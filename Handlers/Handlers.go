package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
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

	currentTime := time.Now()

	w.Write(stuff)

	fmt.Println("Stuff Json accessed at: ", currentTime)
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

	currentTime := time.Now()
	fmt.Println("Boop Json accessed at:", currentTime)

}
