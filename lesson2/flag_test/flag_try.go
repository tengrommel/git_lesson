package main

import (
	"flag"
	"fmt"
	"strings"
)
// String类型的flag
var sep = flag.String("s", " ", "separator")
var newline = flag.Bool("n", false, "append newline")


func main() {
	flag.Parse()
	// *sep 指向内容 Join拼接列表
	fmt.Println(strings.Join(flag.Args(), *sep))
	if *newline {
		fmt.Println()
	}

}

