package main

import (
	 "log"
	 "io/ioutil"
	 "os"
	 "strings"
)

func main() {
	 content, err := ioutil.ReadFile(os.Args[1])
	 if err != nil{
	 	 log.Fatal(err)
	 }
	 count := make(map[string]int)
	 words := strings.Fields(string(content))
	 for _, word :=range words {
			if _, ok := count[word]; ok {
				 count[word]++
			}
	 }
}
