package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	// Define the route handler for "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Log the call
		log.Println("Received request to /")

		// Parse the template file
		tmpl, err := template.ParseFiles("template.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Execute the template and write the output to the response writer
		err = tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// Define the route handler for "/http-get"
	http.HandleFunc("/http-get", func(w http.ResponseWriter, r *http.Request) {
		// Log the call
		log.Println("Received request to /http-get")

		// Extract the URL from the request body
		url := r.FormValue("url")

		// Make a GET request to the specified URL
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Set the Accept header to "application/json"
		req.Header.Set("Accept", "application/json")

		// Send the request and get the response
		client := http.Client{
			Timeout: time.Second * 10, // Timeout after 10 seconds
			// Can be tested with https://httpstat.us/200?sleep=120000
		}
		resp, err := client.Do(req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		// Read the response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Write the response body to the response writer
		w.Write(body)
	})

	// Start the HTTP server
	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
