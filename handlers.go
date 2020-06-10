package main

import (
	"net/http"
	"encoding/json"
	"time"
)

func ping(w http.ResponseWriter, req *http.Request) {
	// https://flaviocopes.com/golang-enable-cors/
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(
		map[string]bool{"pong": false})
}

func blackLivesLost(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		BlackLife{name: "George Floyd", dateOfDeath: time.Now()})
}
