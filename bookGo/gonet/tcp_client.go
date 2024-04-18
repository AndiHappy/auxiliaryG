package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial(ConnType, ConnHost+":"+ConnPort)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ReadNWrite(conn)
	err = conn.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func ReadNWrite(conn net.Conn) {
	message := "Test Request\n"
	_, writeErr := conn.Write([]byte(message))
	if writeErr != nil {
		fmt.Println("failed:", writeErr)
		return
	}
	writeErr = conn.(*net.TCPConn).CloseWrite()
	if writeErr != nil {
		fmt.Println("failed:", writeErr)
		return
	}
	buf, readErr := io.ReadAll(conn)
	if readErr != nil {
		fmt.Println("failed:", readErr)
		return
	}
	fmt.Println(string(buf))
}
