package main

import "fmt"

//全局变量声明
// var 变量名 类型 = 表达式 类型和表达式可以省略其中一个，如果省略类型会根据表达式做类型推到

var boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g  or %g\n", f, c)
}
