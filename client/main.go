package main

import (
	"bufio"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := conn.Close(); err != nil {
			panic(err)
		}
	}()
	print("Enter Your Name: ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	if _, err := conn.Write([]byte(text)); err != nil {
		panic(err)
	}
	println("Your name set to ", text)

	for {
		print("Enter text: ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		if _, err := conn.Write([]byte(text)); err != nil {
			panic(err)
		}
	}

}
