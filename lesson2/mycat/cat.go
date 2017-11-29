package main

import (
	"io/ioutil"
	"fmt"
	"os"
)

func printFile(name string)  {
	buf, err := ioutil.ReadFile(name)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(string(buf))
}

func main() {
	if len(os.Args) != 2{
		fmt.Println("打印哪个文件？")
		return
	}
	printFile(os.Args[1])
}
