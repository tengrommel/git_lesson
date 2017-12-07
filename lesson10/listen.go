package main

import (
	"net"
	"flag"
	"io"
	"log"
	"sync"
)

var (
	target = flag.String("target", "www.qq.com:80", "target host")
)

func handleConn(conn net.Conn)  {
	var remote net.Conn
	var err error
	remote, err = net.Dial("tcp", *target)
	if err!=nil{
		log.Print(err)
		conn.Close()
		return
	}
	wg := new(sync.WaitGroup)
	wg.Add(2)
	// go 读取(conn)的数据，发送到remote, 直到conn的EOF，关闭remote
	go func() {
		defer wg.Done()
		io.Copy(remote, conn)
		remote.Close()
	}()
	// go 读取remote的数据，发送到客户端(conn)，直到remote的EOF，关闭conn
	go func() {
		defer wg.Done()
		io.Copy(conn, remote)
		conn.Close()
	}()
	wg.Wait()
	// 等待两个协程结束
}

func main() {
	// 建立listen
	l, err := net.Listen("tcp", ":8021")
	if err != nil{
		log.Fatal(err)
	}
	for {
		conn, _ := l.Accept()
		go handleConn(conn)
	}

}
