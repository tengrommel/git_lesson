package main

import (
	 "fmt"
	 "unsafe"
)

func main() {
	 var (
	 	 x byte
	 	 x1 int // 32位平台和64位平台不一样 32位为4 64位为8
	 	 x2 int32
	 	 x3 int64
	 	 x4 uint
	 	 x5 uint32
	 	 x6 uint64
	 )
	fmt.Println(x, x1, x2, x3, x4, x5, x6)
	fmt.Println(unsafe.Sizeof(x1))
	str1 := "hello"
	doc := `你好
		我换行了`
	fmt.Println(str1, doc)
	var b byte
	for b=0; b<177; b++{
		 fmt.Printf("%d %c\n", b, b)
	}
}
