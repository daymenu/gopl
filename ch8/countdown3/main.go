package main

import (
	"fmt"
	"os"
	"time"
)

var j time.Time

func main() {
	fmt.Println("Commencing countdown")
	ticker := time.NewTicker(1 * time.Second)

	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	for countdown := 10; countdown > 0; countdown-- {
		select {
		case <-ticker.C:
			fmt.Println(countdown)
		case <-abort:
			fmt.Println("数据异常，停止发射")
			os.Exit(2)
		}
	}
	ticker.Stop()
	lunach()
}

func lunach() {
	height := 0
	tick := time.Tick(1 * time.Second)
	for i := 0; ; i++ {
		<-tick
		if height < 30000 {
			height += i * i
			fmt.Printf("火箭高度%dkm\n", height)
		} else {
			fmt.Printf("环球飞行-火箭高度%dkm\n", height)
		}

	}
}
