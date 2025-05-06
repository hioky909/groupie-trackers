package main

import (
	"fmt"
	handler "groupie-tracker/handlers"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/artist/", handler.DetailsHandler)
	mux.HandleFunc("/home", handler.HomeHandler)
	fmt.Println("Server launching on port 8080.")
	fmt.Println("Server is launched.")
	fmt.Println("Link: http://localhost:8080/home")
	server := &http.Server{
		Addr:              ":8080",           //adresse du server (le port choisi est à titre d'exemple)
		Handler:           mux,               // listes des handlers
		ReadHeaderTimeout: 10 * time.Second,  // temps autorisé pour lire les headers
		WriteTimeout:      10 * time.Second,  // temps maximum d'écriture de la réponse
		IdleTimeout:       120 * time.Second, // temps maximum entre deux rêquetes
		MaxHeaderBytes:    1 << 20,           // 1 MB // maxinmum de bytes que le serveur va lire
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/artist/", handler.DetailsHandler)
	http.HandleFunc("/artist/", handler.LocationHandler)
}
