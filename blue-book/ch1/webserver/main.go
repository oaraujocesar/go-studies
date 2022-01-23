package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	lissajous "github.com/oaraujocesar/go-studies/blue-book/ch1/gifs"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", countHandler)
	http.HandleFunc("/lissajous", lissajousHandler)

	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()

	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)

	for key, value := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", key, value)
	}

	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for key, value := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", key, value)
	}

}

func countHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()

	fmt.Fprintf(w, "Count %d\n", count)

	mu.Unlock()
}

func lissajousHandler(w http.ResponseWriter, r *http.Request) {

	qc := r.URL.Query()["cycles"][0]

	cycles, _ := strconv.Atoi(qc)

	lissajous.Lissajous(w, cycles)
}
