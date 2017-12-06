package main

import "fmt"

func fibonacci_(c chan int, quit chan bool)  {
	x, y := 0, 1
	for  {
		select{
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- true
	}()
	fibonacci_(c, quit)
}