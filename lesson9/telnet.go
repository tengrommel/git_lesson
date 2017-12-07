package main

import (
	"net"
	"log"
	"fmt"
	"io"
	"os"
)

/*
文件 http.RespBody chan socket 需要关闭
 */
func main() {
	addr := "127.0.0.1:8021"
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
	//buf := make([]byte, 4096)
	//n, err =conn.Read(buf)
	//if err != nil && err != io.EOF{
	//	log.Fatal(err)
	//}
	//fmt.Println(string(buf[:n]))
	// 实现按行读取
	io.Copy(os.Stdout, conn)
	conn.Close()
}
