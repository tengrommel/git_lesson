package main

import "fmt"

func ol() {
	defer func() {
		fmt.Println("defer")
	}()
	fmt.Println("hello")
}

func main()  {
	ol()
}

