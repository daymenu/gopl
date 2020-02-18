package main

import (
	"fmt"
	"log"
	"os"
)

type People struct {
	name string
	age  int
	sex  int
}
type Emplyee struct {
	People
	title string // 职称
}

type Speaker interface {
	Speak()
}

const boilingF = 212.0

func (p *People) Speak() {
	fmt.Println(p.name)
}

func Talk(s Speaker) {
	s.Speak()
}
func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F  or %g°C\n", f, c)
	p := People{name: "lili"}
	Talk(&p)

	log.SetOutput(os.Stdout)
	log.Println("hello")

	var a interface{}
	a = "aa"
	switch a.(type) {
	case int:
		fmt.Println(" i am int")
	case string:
		fmt.Println("i am string")
	}
}
