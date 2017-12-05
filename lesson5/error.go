package main

import (
	"os"
	"io/ioutil"
	"log"
	"time"
	"fmt"
)

func read(f *os.File) (string, error) {
	buf, err := ioutil.ReadAll(f)
	if err != nil{
		return "", err
	}
	return string(buf), nil
}

func main() {
	f, err := os.Open("a.txt")
	if err != nil{
		log.Fatal(err)
	}
	defer f.Close()
	var content string
	retries := 3
	for i:=1;i<=retries;i++ {
		content, err = read(f)
		if err == nil{
			break
		}
		time.Sleep(time.Second)
	}
	fmt.Println(content)
}
