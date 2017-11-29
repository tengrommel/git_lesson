package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	fmt.Println("first", os.Args[1])
	fmt.Println("second", os.Args[2])
}
