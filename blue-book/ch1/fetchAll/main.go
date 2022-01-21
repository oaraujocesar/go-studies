package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch) // starts a goroutine
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) // receives from channel
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()

	res, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err) // send the erro to the channel

		return
	}

	nbytes, err := io.Copy(ioutil.Discard, res.Body)

	defer res.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)

		return
	}

	secs := time.Since(start).Seconds()

	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
