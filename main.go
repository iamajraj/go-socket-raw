package main

import (
	"fmt"
	"net"
)

func main() {
	server, err := net.Listen("tcp", "127.0.0.1:8000")
	defer func() {
		if err := server.Close(); err != nil {
			panic(err)
		}
	}()

	println("listening on port :8000")

	if err != nil {
		panic(err)
	}

	for {
		conn, err := server.Accept()
		if err != nil {
			panic(err)
		}
		var name = make([]byte, 1024)
		len, err := conn.Read(name)
		if err != nil {
			panic(err)
		}
		go handle(conn, string(name[:len]))
	}

}

func handle(conn net.Conn, name string) {
	conn.Write([]byte("Hello From Server"))
	var msg = make([]byte, 1024)
	for {
		mlen, err := conn.Read(msg)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Msg of %s: %s", name, string(msg[:mlen]))
	}
}
