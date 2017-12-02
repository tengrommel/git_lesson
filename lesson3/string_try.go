package main

import (
	 "fmt"
)


func main() {
	 str1 := "hello\\world\b" // 制表符 \b
	 str2 := ""
	 if str2 == str1{

	 }
	 // byte
	 c1 := str1[1]
	 fmt.Println(0, len(str1)-1)
	 fmt.Println(str1)
	 fmt.Println(c1)

	 s3 := str1[3:6]
	 fmt.Println(s3)
}
