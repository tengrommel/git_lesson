package main

import (
	"net"
	"log"
	"os"
	"io/ioutil"
	"bufio"
	"strings"
)

func handleConn(conn net.Conn)  {
	defer conn.Close()
	r := bufio.NewReader(conn)
	line, _ := r.ReadString('\n')
	line = strings.TrimSpace(line)
	fields := strings.Fields(line)
	if len(fields)!=2{
		conn.Write([]byte("bad input"))
		return
	}
	cmd := fields[0]
	name := fields[1]
	if cmd == "GET" {
		f, err := os.Open(name)
		if err != nil {
			log.Fatal(err)
			return
		}
		defer f.Close()
		buf, err := ioutil.ReadAll(f)
		if err != nil {
			log.Print(err)
			return
		}
		conn.Write(buf)
	}else if cmd == "STORE"{
		// 从r读取文件内容
	}
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
