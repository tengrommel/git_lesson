package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("netstat", "-an")
	out, err := cmd.CombinedOutput()
	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println(string(out))
}
