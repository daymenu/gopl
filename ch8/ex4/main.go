package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration, wg sync.WaitGroup) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
	time.Sleep(delay)
	wg.Done()
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	var wg sync.WaitGroup
	for input.Scan() {
		wg.Add(1)
		go echo(c, input.Text(), 1*time.Second, wg)
	}

	go func() {
		wg.Wait()
		c.CloseWrite()
	}()
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
		fmt.Printf("%v is connected\n", conn.RemoteAddr())
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
