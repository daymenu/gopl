package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
)

func main() {
	port := flag.Int("port", 21, "ftp 端口号")
	flag.Parse()
	address := fmt.Sprintf("localhost:%d", *port)
	fmt.Println("ftp is started: ", address)
	linstener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	for {
		client, err := linstener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go echo(client)
	}
}

func echo(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		fmt.Println(input.Text())
	}
	c.Close()
}
