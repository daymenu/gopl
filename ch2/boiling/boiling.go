package main

import (
	"fmt"
)

const boilingF = 212.0

const (
	//Monday 星期一
	Monday = 2 + iota
	//Tuesday 星期二
	Tuesday
	//Wednesday 星期三
	Wednesday
	//Thursday 星期四
	Thursday
	//Friday 星期五
	Friday
	//Saturday 星期六
	Saturday
	//Sunday 星期日
	Sunday
)

const (
	_ = 1 << (10 * iota)
	KiB
	MiB
	GiB
	TiB
	PiB
	EiB
	ZiB
	YiB
)

const (
	s1 = 1
	s2
	s3
)

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F  or %g°C\n", f, c)
}
