package main

import "fmt"

// 切片
func sum(args ...int) int {
	n := 0
	for i:=0;i<len(args);i++ {
		n += args[i]
	}
	return n
}

func main() {
	fmt.Println(sum(1,2,3))
}
