package main

import (
	"fmt"
	"log"
	"os"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT");
	if port == "" {
		port = "8080"
	}

	r := mux.NewRouter()
	r.HandleFunc("/", ping)
	r.HandleFunc("/blackLivesLost", blackLivesLost)

	http.Handle("/", r)
	fmt.Println("Starting up on " + port)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}
