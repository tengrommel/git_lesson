package main

import (
	 "os"
	 "log"
	 "fmt"
)

func main() {
	 f, err := os.Open(".")
	 if err != nil{
	 	 log.Fatal(err)
	 }
	 infos, err := f.Readdir(-1)
	 if err != nil{
			log.Fatal(err)
	 }
	 for _, info := range infos{
	 	 fmt.Printf("%v %d %s\n", info.IsDir(), info.Size(), info.Name())
	 }
}
