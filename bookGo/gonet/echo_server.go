package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

const readAnyBufSize = 4096

// readAny returns a byte slice with data read from a reader
// as long as read length is less than readAnyBufSize,
// the byte slice length will match that of the read length
func readAny(r io.Reader) ([]byte, error) {
	buf := make([]byte, readAnyBufSize)
	n, err := r.Read(buf)
	return buf[:n], err
}

func client(conn net.Conn) {
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	userInput := bufio.NewReader(os.Stdin)
	for {
		// get user input
		fmt.Print("enter a string: ")
		s, err := userInput.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		// send user input to echo server
		n, err := io.WriteString(conn, s[:len(s)-1])
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Send %v bytes", n)
		if n == 0 {
			continue // no point waiting for a reply if we didn't send any data
		}
		// receive reply from echo server
		buf, err := readAny(conn)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Receive (%v bytes): `%v`", len(buf), string(buf))
	}
}

func server(conn net.Conn) {
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	for {
		// receive message from client
		buf, err := readAny(conn)
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

func serve(ln net.Listener) {
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Accept connection from", conn.RemoteAddr())
		go server(conn)
	}
}

func main() {
	var (
		address  = flag.String("addr", ":9989", "listen/dial address")
		isServer = flag.Bool("serve", true, "is this a client or server?")
	)
	flag.Parse()
	if *isServer {
		ln, err := net.Listen("tcp", *address)
		if err != nil {
			log.Fatal(err)
		}
		serve(ln)
	} else {
		conn, err := net.Dial("tcp", *address)
		if err != nil {
			log.Fatal(err)
		}
		client(conn)
	}
}
