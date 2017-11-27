package main

import "fmt"

func main() {
	var x int
	x = 1
	y := 3
	var p *int
	p = &x
	fmt.Println(p) // 指针里存的就是地址
	fmt.Println(*p) // *取内容
	fmt.Println("x=",x) // *取内容
	*p = 2				// 修改指针的值
	add(&x)
	fmt.Println("x=",x)
	p = &y
	fmt.Println("y=", *p)
	fmt.Println("hello golang")
}

func add(p *int)  {
	*p = *p + 2
}