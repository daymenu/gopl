package main

import (
	"fmt"
	"io/ioutil"
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
			log.Printf("fetch %v\n", err)
		}
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("read data %s is faild, error %v", url, err)
		}
		fmt.Println(string(data))
	}
}

func httpPrefix(url string) string {
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	return url
}
