package main

import (
	"net"
	"log"
	"os"
	"io/ioutil"
)

func handleConn(conn net.Conn)  {
	defer conn.Close()
	f, err := os.Open("index.html")
	if err != nil{
		log.Fatal(err)
		return
	}
	defer f.Close()
	buf, err := ioutil.ReadAll(f)
	if err != nil{
		log.Print(err)
		return
	}
	conn.Write(buf)
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
