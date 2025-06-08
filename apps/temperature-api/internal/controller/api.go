package controller

import (
	"log"
	"net/http"
)

func StartServer(appAddress string) {
	mux := http.NewServeMux()

	mux.HandleFunc("/temperature", randomTemperature)
	mux.HandleFunc("/temperature/{id}", randomTemperature)

	log.Fatal(http.ListenAndServe(appAddress, mux))
}
