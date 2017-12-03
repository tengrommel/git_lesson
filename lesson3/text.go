package main

import (
	 "os"
	 "log"
)

func main() {
	 f, err := os.Create("a.txt")
	 if err != nil{
	 	 log.Fatal(err)
	 }
	 f.WriteString("hello\n")
	 f.Close()
}
