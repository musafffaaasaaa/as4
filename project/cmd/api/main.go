package main

import (
	"net/http"
	"project/handler"
)

func main() {
	// Create a file server to serve static assets (e.g., CSS, JavaScript, images)
	fs := http.FileServer(http.Dir("static"))

	// Serve static assets under the "/static/" URL prefix
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Initialize your web handlers
	handler.NewHandler()

	// Define your HTTP routes
	http.HandleFunc("/shop/add", func(w http.ResponseWriter, r *http.Request) {
		// Implement your "/shop/add" handler logic here
	})

	http.HandleFunc("/shop/get", func(w http.ResponseWriter, r *http.Request) {
		// Implement your "/shop/get" handler logic here
	})

	http.HandleFunc("/shop/delete", func(w http.ResponseWriter, r *http.Request) {
		// Implement your "/shop/delete" handler logic here
	})

	http.HandleFunc("/shop/drop", func(w http.ResponseWriter, r *http.Request) {
		// Implement your "/shop/drop" handler logic here
	})

	// Start the HTTP server on port 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err) // Handle server startup error if any
	}
}
