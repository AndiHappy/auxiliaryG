package main

import (
	"io"
	"log"
	"net"
)

func readNotBlocked(r io.Reader) ([]byte, error) {
	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	return buf[:n], err
}

func handleConn(conn net.Conn) {
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	for {
		// receive message from client
		buf, err := readNotBlocked(conn)
		if err != nil {
			log.Print(err)
			return
		}
		log.Printf("Receive (%v bytes): `%v`", len(buf), string(buf))
		// reply with same message
		n, err := conn.Write(buf)
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("Send (%v bytes)", n)
	}
}

func handleListen(ln net.Listener) {
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Accept connection from", conn.RemoteAddr())
		go handleConn(conn)
	}
}

func main() {
	ln, err := net.Listen("tcp", "localhost:9988")
	if err != nil {
		log.Fatal(err)
	}
	handleListen(ln)
}
