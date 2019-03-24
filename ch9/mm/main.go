package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gopl/ch9/memo2"
)

func main() {
	data := []string{
		// "http://www.baidu.com/",
		// "http://www.baidu.com/",
		"http://www.xgo.cn/",
		"http://www.xgo.cn/",
	}

	m := memo2.New(httpGetBody)
	for _, url := range data {
		url := url
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Println(err)
		}
		fmt.Printf("%s %ds %d bytes", url, time.Since(start), len(value.([]byte)))
	}
}

func Hello(url string) (interface{}, error) {
	return []byte(url), nil
}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func incomingURLs() []string {
	return []string{"http://www.baidu.com",
		"http://www.xgo.cn",
		"http://www.baidu.com",
		"http://www.xgo.cn"}
}
