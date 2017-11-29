package main

import "fmt"

func main() {
	//var x int
	var p *int
	// lesson2/point.go:9: cannot use 291 (type int) as type *int in assignment
	// 不能将数字赋值给指针
	p = 0x0123
	fmt.Println(*p)
}
