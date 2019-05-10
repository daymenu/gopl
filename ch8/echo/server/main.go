package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"sync"
)

type Online struct {
	Num int
	sync.Mutex
}

var online Online

func main() {
	listen, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("server is start ")
	for {
		conn, err := listen.Accept()
		online.Lock()
		online.Num += 1
		fmt.Println(conn.LocalAddr(), "is connected,total:", online.Num)
		online.Unlock()

		if err != nil {
			log.Println(err)
		}
		go handelConn(conn)
	}
}

func handelConn(conn net.Conn) {
	defer func() {
		online.Lock()
		online.Num += 1
		online.Unlock()
		conn.Close()
	}()
	input := bufio.NewScanner(conn)
	for {
		if input.Scan() {
			text := input.Text()
			fmt.Fprintf(conn, "server say :%s\n", text)
			fmt.Println(text)
		}
	}
}
