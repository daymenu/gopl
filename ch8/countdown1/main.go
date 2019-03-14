package main

import (
	"fmt"
	"time"
)

var j time.Time

func main() {
	fmt.Println("Commencing countdown")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		j = <-tick
	}
	lunach()
}

func lunach() {
	for {
		fmt.Println(time.Now())
	}
}
