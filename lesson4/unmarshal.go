package main

import (
	"encoding/json"
	"log"
	"fmt"
)

type StudentStruct struct {
	Id int
	Name string
}

func main() {
	str := `{"Id":2, "Name":"alice"}`
	var s StudentStruct
	err := json.Unmarshal([]byte(str), &s)
	if err != nil{
		log.Fatal("unmarshal error:%s", err)
	}
	fmt.Println(s)
}
