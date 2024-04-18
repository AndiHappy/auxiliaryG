package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

const (
	ConnHost = "localhost"
	ConnPort = "3333"
	ConnType = "tcp"
)

func main() {
	l, err := net.Listen(ConnType, ConnHost+":"+ConnPort)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer func() {
		err := l.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	fmt.Println("Listening on " + ConnHost + ":" + ConnPort)

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	buf, readErr := io.ReadAll(conn)
	if readErr != nil {
		fmt.Println("failed:", readErr)
		return
	}
	fmt.Println("Got: ", string(buf))

	_, writeErr := conn.Write([]byte("Message received.\n"))
	if writeErr != nil {
		fmt.Println("failed:", writeErr)
		return
	}
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
}
