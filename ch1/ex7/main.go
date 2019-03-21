package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		url = httpPrefix(url)
		resp, err := http.Get(url)
		if err != nil {
			log.Printf("fetch: %v\n", err)
		}
		defer resp.Body.Close()
		_, err = io.Copy(os.Stdin, resp.Body)
		if err != nil {
			log.Printf("read: %v\n", err)
		}
	}
}

func httpPrefix(url string) string {
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}

	return url
}
