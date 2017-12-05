package main

import (
	"sort"
	"fmt"
)

func main() {
	s := []int{2,3,1,5,9,7}
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	fmt.Println(s)
}
