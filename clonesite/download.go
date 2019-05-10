package main

import (
	"io"
	"net/http"
	"os"
)

func download(url string, path string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	f, err := os.Create(path)
	if err != nil {
		return err
	}

	defer f.Close()
	io.Copy(f, resp.Body)
	return err
}
