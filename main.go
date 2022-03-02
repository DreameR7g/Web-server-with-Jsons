package main

import (
	"fmt"
	"log"
	"net/http"
	"web-server-with-jsons/handlers"
)

func main() {

	fileServer := http.FileServer(http.Dir("./"))
	http.Handle("/", fileServer)
	http.HandleFunc("/json", handlers.JsonHandler)
	http.HandleFunc("/boop", handlers.BoopHandler)

	fmt.Printf("Starting server at port 1338\n")
	if err := http.ListenAndServe(":1338", nil); err != nil {
		log.Fatal(err)
	}

}
