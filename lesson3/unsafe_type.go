package main

import (
	 "fmt"
	 "unsafe"
)

func main() {
	 var (
	 	 x byte
	 	 x1 int
	 	 x2 int32
	 	 x3 int64
	 	 x4 uint
	 	 x5 uint32
	 	 x6 uint64
	 )
	fmt.Println(x, x1, x2, x3, x4, x5, x6)
	fmt.Println(unsafe.Sizeof(x1))
}
