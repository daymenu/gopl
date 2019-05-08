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
	ch := make(chan string) //由于该chan没有缓存，所以go会在 39行暂停处暂停
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	time.Sleep(1 * time.Second)
}

func fetch(url string, ch chan<- string) {
	fmt.Println("fetch:", url)
	start := time.Now()
	url = httpPrefix(url)
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("fetch: %v\n", err)
	}
	defer resp.Body.Close()
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		log.Printf("read: %v\n", err)
	}
	ch <- fmt.Sprintf("%.2fs %7d %s", time.Since(start).Seconds(), nbytes, url) //暂停处
	fmt.Println("end fetch", url)
}

func httpPrefix(url string) string {
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}

	return url
}
