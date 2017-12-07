package main

import (
	"net"
	"log"
)

func main() {
	addr := ":8021"
	listener, err := net.Listen("tcp", addr)
	if err != nil{
		log.Fatal(err)
	}
	defer listener.Close()
	conn, err := listener.Accept()
	if err!= nil{
		log.Fatal(err)
	}
	conn.Write([]byte("hello golang!\n"))
	conn.Close()
}
