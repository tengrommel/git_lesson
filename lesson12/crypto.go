package main

import (
	"io"
	"net"
)

type CryptoWriter struct {
	w io.Writer
}

func NewCryptoWriter(w io.Writer) io.Writer {
	return &CryptoWriter{
		w: w,
	}
}

func (w *CryptoWriter)Write(b []byte) (int, error) {
	return 1, nil
}

type CryptoReader struct {
	r io.Reader
}

func NewCryptoReader(r io.Reader) io.Reader {
	return &CryptoReader{
		r: r,
	}
}

func (r *CryptoReader)Read(b []byte) (int, error) {
	return 1, nil
}

func main() {
	remote, err := net.Dial("tcp", "")
	w := NewCryptoWriter(remote)
	w.Write([]byte("hello"))
	r := NewCryptoReader(remote)
	buf := make([]byte, 1024)
	r.Read(buf)
	if err != nil{

	}
}
