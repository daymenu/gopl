package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func main() {
	fmt.Println(fetch("http://www.baidu.com/"))
}
func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, fmt.Errorf("get %s:%v", url, err)
	}
	defer resp.Body.Close()
	local := path.Base(resp.Request.URL.Path)
	if local == "/" || local == "." {
		local = "index.html"
	}
	fmt.Println(local)
	f, err := os.Create(local)
	if err != nil {
		return "", 0, fmt.Errorf("create %s:%v", local, err)
	}
	n, err = io.Copy(f, resp.Body)
	// Close file, but prefer error from Copy,if any.
	if closeErr := f.Close(); err != nil {
		err = closeErr
	}
	return local, n, err
}
