package main

import "fmt"

const PoundRate = 0.4535924

type Pound float64
type Kilo float64

// 磅转化为公斤
func PToK(p Pound) Kilo {
	return Kilo(p * PoundRate)
}

// 公斤转化为磅
func KToP(k Kilo) Pound {
	return Pound(k / PoundRate)
}

//磅格式化输出
func (p Pound) String() string {
	return fmt.Sprintf("%f lb", p)
}

//公斤格式化输出
func (k Kilo) String() string {
	return fmt.Sprintf("%f kg", k)
}
