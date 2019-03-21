package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
}

func fetch(url string, ch chan<- string) {
	url = httpPrefix(url)
	start := time.Now()
	resp, err := http.Get(url)
	oneSec := time.Since(start).Seconds()
	start1 := time.Now()
	resp, err = http.Get(url)
	twoSec := time.Since(start1).Seconds()
	if err != nil {
		log.Printf("fetch: %v\n", err)
	}
	defer resp.Body.Close()
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		log.Printf("read: %v\n", err)
	}
	ch <- fmt.Sprintf("%.2fs %.2fs %7d %s", oneSec, twoSec, nbytes, url)
}

func httpPrefix(url string) string {
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}

	return url
}
