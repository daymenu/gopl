package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	lists := os.Args[1:]
	num := len(os.Args[1:])
	walls := make(map[string]string, num)
	ch := make(chan int, num)
	for _, l := range lists {
		areas := strings.Split(l, "=")
		conn, err := net.Dial("tcp", areas[1])
		if err != nil {
			log.Fatal(err)
		}
		go func() {
			scan := bufio.NewScanner(conn)
			if scan.Scan() {
				walls[areas[0]] = scan.Text()
				ch <- 1
			}
		}()
		defer conn.Close()
	}
	for i := 0; i < num; i++ {
		<-ch
	}
	for k, time := range walls {
		fmt.Printf("%s:%s\n", k, time)
	}
}
