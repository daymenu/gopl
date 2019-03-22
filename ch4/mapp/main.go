package main

import "fmt"

func main() {
	stu := map[string]int{
		"yuwen":  90,
		"shuxue": 80,
	}
	if score, ok := stu["kexue"]; !ok {
		fmt.Println("kexue is not exists", score)
	}
	delete(stu, "shuxue")
	fmt.Println(stu)
}
