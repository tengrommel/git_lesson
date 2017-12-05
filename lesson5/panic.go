package main

import "fmt"

func pp()  {
	var p *int
	fmt.Println(*p)
}

func main() {
	defer func() {
		err := recover()
		fmt.Println(err)
	}()
	pp()
}
