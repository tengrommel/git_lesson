package main

import (
	"fmt"
	"strings"
)

func PrintExample() int {
	return 10
}

func smap(r rune) rune {
	fmt.Printf("%c\n", r)
	return r
}

func main() {
	var f func() int
	// 匿名函数
	f = func() int {
		return 10
	}
	fmt.Println(f())
	fmt.Println()
	s := strings.Map(func(r rune) rune {
		return r-32
	}, "hello")
	fmt.Println(s)
}
