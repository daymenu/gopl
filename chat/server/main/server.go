package main

import (
	"fmt"
	"net"
)

func runServer(addr string) (err error) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("listen failed, ", err)
		return
	}
	fmt.Println("server is start at :", addr)
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("accept failed, ", err)
			continue
		}
		go proccess(conn)
	}
}

func proccess(s net.Conn) {
	defer s.Close()
}
