package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	m := http.NewServeMux()
	m.HandleFunc("/", handlePage)

	// Read port from environment variable, fallback to 8010
	port := os.Getenv("PORT")
	if port == "" {
		port = "8010"
	}

	srv := http.Server{
		Handler:      m,
		Addr:         ":" + port,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	fmt.Println("server started on ", port)
	err := srv.ListenAndServe()
	log.Fatal(err)
}
