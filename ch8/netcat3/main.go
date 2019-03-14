package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{}, 1)
	go func() {
		log.Println("recieved data")
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}
	}()

	go mustCopy(conn, os.Stdin)
	<-done
	conn.Close()
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
