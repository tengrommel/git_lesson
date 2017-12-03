package main

import (
	 "encoding/json"
	"log"
	"fmt"
)

type StudentExample struct {
	 Id int
	 Name string
}


func main() {
	s := StudentExample{
		Id: 2,
		Name: "alice",
		}
	buf, err := json.Marshal(s)
	if err != nil{
		log.Fatal("Marshal error:%s", err)
	}
	fmt.Println(string(buf))
}
