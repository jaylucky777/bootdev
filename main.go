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

	// Handle all requests to "/"
	m.HandleFunc("/", handlePage)

	// Read port from environment variable, fallback to 8010
	port := os.Getenv("PORT")
	if port == "" {
		port = "8010" // default port if not set
	}

	// Configure the HTTP server
	srv := http.Server{
		Handler:      m,
		Addr:         ":" + port,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	// Start the server
	fmt.Println("Server started on port", port)
	err := srv.ListenAndServe()
	log.Fatal(err)
}

// handlePage writes a simple HTML page
func handlePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	const page = `<html>
<head></head>
<body>
	<p>Hello from Docker! I'm a Go server.</p>
</body>
</html>`
	w.Write([]byte(page))
}
