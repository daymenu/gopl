package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	client, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		log.Fatal(err)
	}
	go mustCopy(os.Stdout, client)
	mustCopy(client, os.Stdin)
	client.Close()
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
