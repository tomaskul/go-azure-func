package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleRoute)
}

func handleRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "My not-so-first REST API")
}
