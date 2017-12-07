package main

import (
	"net"
	"log"
	"fmt"
)

/*
文件 http.RespBody chan socket 需要关闭
 */
func main() {
	addr := "123.125.114.144:80"
	conn, err := net.Dial("tcp", addr)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(conn.RemoteAddr().String())
	fmt.Println(conn.LocalAddr().String())
	n, err := conn.Write([]byte("GET / HTTP/1.1\r\n\r\n"))
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("write size", n)
	buf := make([]byte, 4096)
	n, err =conn.Read(buf)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(string(buf[:n]))
	conn.Close()
}
