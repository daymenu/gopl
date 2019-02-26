package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
	time.Sleep(delay)
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)
	}
	c.Close()
}

func main() {
	port := flag.Int("port", 8000, "请输入要监听的端口号")
	flag.Parse()

	address := fmt.Sprintf("localhost:%d", *port)
	fmt.Printf("reverb1 is started at : %s\n", address)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		fmt.Printf("%v is connected", conn.RemoteAddr())
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
