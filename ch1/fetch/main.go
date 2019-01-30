package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		httpPrefix := "http://"
		if !strings.HasPrefix(url, httpPrefix) {
			url = httpPrefix + url
		}
		res, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(res.Body)
		// io.Copy(os.Stdout, res.Body) 练习1.7
		res.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: read %s %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%d\n%s", res.StatusCode, string(b))
	}
}
