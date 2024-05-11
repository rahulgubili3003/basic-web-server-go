package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/basic", handler)
	// Listen on port 8080
	fmt.Println("Server is listening on port 8085...")
	if err := http.ListenAndServe(":8085", nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Welcome to the basic Http server")
	if err != nil {
		return
	}
}
