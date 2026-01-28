package main

import (
	"log"
	"net/http"
)

func main() {
	const filepath = "../client"
	const port = "8080"
	mux := http.NewServeMux()
	mux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir(filepath))))

	mux.HandleFunc("/healthz", handlerReadiness)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Server is running on port:%s", port)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
