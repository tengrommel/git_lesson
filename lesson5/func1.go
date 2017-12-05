package main

import "fmt"

func printHello()  {
	fmt.Println("hello")
}

func printHello1()  {
	fmt.Println("hello1")
}

func func1(n int) int {
	return n+1
}

type fstruct struct {
	Func func()
} 

func main() {
	var ok func()
	ok = printHello
	ok()
	var flist [3]func(int) int
	var fslice []func()
	var fmap map[string]func()

	flist[0] = func1
	flist[0](10)
	fslice[0]= printHello
	fslice[0]()
	fmap["this"] = printHello
	fmap["this"]()
}