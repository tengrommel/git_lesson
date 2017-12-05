package main

import "fmt"

func iter(s []int) func() int {
	var i = 0;
	return func() int {
		if i >= len(s) {
			return 0
		}
		n := s[i]
		i+=1
		return n
	}
}

func main() {
	f := iter([]int{1,2,3})
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
