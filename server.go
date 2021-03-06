package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("Server running on port: %d", 8080)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "My name is %s", "Max")
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
