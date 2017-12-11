package main

import (
	"io"
	"net"
	"crypto/rc4"
)

type CryptoWriter struct {
	w io.Writer
	cipher *rc4.Cipher
}

func NewCryptoWriter(w io.Writer, key string) io.Writer {
	return &CryptoWriter{
		w: w,
	}
}

// 把b里面的数据进行加密，之后写入到w.w里面
// 调用w.w.Write进行写入
func (w *CryptoWriter)Write(b []byte) (int, error) {
	return 1, nil
}

type CryptoReader struct {
	r io.Reader
}

func NewCryptoReader(r io.Reader, key string) io.Reader {
	return &CryptoReader{
		r: r,
	}
}

func (r *CryptoReader)Read(b []byte) (int, error) {
	return 1, nil
}

func main() {
	remote, err := net.Dial("tcp", "")
	w := NewCryptoWriter(remote,"123456")
	w.Write([]byte("hello"))
	r := NewCryptoReader(remote, "123456")
	buf := make([]byte, 1024)
	r.Read(buf)
	if err != nil{

	}
}
