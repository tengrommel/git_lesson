package main

import (
	"os"
	"log"
	"os/exec"
)

func main() {
	f, err := os.Create("netstat.out")
	if err !=nil{
		log.Fatal(err)
	}
	defer f.Close()
	cmd := exec.Command("netstat", "-an")
	cmd.Stdout = f
	cmd.Start()
	cmd.Wait()
}
