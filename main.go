package main

import (
	handlers "hng-task-one/api"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/api/hello", handlers.HelloHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Server is running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
