package main

import "fmt"

func main() {
	name := "小明"
	p := &name

	fmt.Printf("%s %[1]T %#x\n", name, &name)
	fmt.Printf("%s %[1]T %#x %#x\n", *p, p, &p)

	num := new(int)
	*num = 3
	p1 := num
	fmt.Printf("%d %#x %#x\n", *num, num, &num)
	fmt.Printf("%d %#x %#x\n", *p1, p1, &p1)
}
