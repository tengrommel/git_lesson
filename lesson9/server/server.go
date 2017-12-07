package main

import (
	"net"
	"log"
)

func handleConn(conn net.Conn)  {
	conn.Write([]byte("hello golang!\n"))
	conn.Close()
}

func main() {
	addr := ":8021"
	listener, err := net.Listen("tcp", addr)
	if err != nil{
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConn(conn)
	}
}
