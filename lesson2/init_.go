package main

import "fmt"

/*
初值
bool值的默认值为false
 */

func main() {
	var (
		x int
		y float32
		z string
		p *int
		boolValue bool
	)
	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%v\n", z)
	fmt.Printf("%v\n", p)
	fmt.Printf("%v\n", boolValue)
}
