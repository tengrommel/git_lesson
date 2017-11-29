package main

import "fmt"

/*
初值
bool值的默认值为false
初始化为零值 安全
 */
var (
		x int
		y float32
		z string
		p *int
		boolValue bool
	)

func main() {
	i:=0
	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%v\n", z)
	fmt.Printf("%v\n", p)
	fmt.Printf("%v\n", boolValue)
	fmt.Printf("%v\n", i)
}
