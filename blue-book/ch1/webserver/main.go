package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", countHandler)

	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()

	fmt.Fprintf(w, "URL.path = %q", r.URL.Path)
}

func countHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()

	fmt.Fprintf(w, "Count %d\n", count)

	mu.Unlock()
}
