package main

import (
	"fmt"
	"time"
)

// ShowTime 更易懂的时间
const ShowTime = "2006-01-02 15:04:05"

func counter(naturals chan<- int) {
	for i := 0; i < 10; i++ {
		naturals <- i
	}
	close(naturals)
}

func squarer(squares chan<- int, naturals <-chan int) {
	for i := range naturals {
		squares <- i * i
	}
	close(squares)
}

func printer(squares <-chan int) {
	for v := range squares {
		fmt.Println(v)
	}
}

func main() {
	timer := time.Tick(time.Second)
loop:
	for i := 0; i < 10; i++ {
		for j := i; j < 20; j++ {
			fmt.Println(<-timer)
			if j == 15 {
				fmt.Println("loop")
				break loop
			}
		}
	}

	newTimer := time.NewTicker(time.Second)
	<-newTimer.C
	newTimer.Stop()
	nowTime := time.Now()
	t := nowTime.Format(ShowTime)

	tenAfter := nowTime.Add(10 * time.Minute)

	d := tenAfter.Sub(nowTime)

	fmt.Println(d)

	fmt.Println(t)

	fmt.Println(tenAfter.Format(ShowTime))

	fmt.Println(time.Parse(ShowTime, tenAfter.Format(ShowTime)))
}
