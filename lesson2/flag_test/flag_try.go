package main

import (
	"flag"
	"fmt"
	"strings"
)
// String类型的flag
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	// *sep 指向内容
	fmt.Println(strings.Join(flag.Args(), *sep))
}

