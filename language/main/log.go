package main

import (
	"log"
	"net/smtp"
)

func init() {
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("init started")
}
func Test1() {
	// Println writes to the standard logger.
	log.Println("main started")
	// Fatalln is Println() followed by a call to os.Exit(1)
	log.Fatalln("fatal message")
	// Panicln is Println() followed by a call to panic()
	log.Panicln("panic message")
}

// Program in GO language with real world example of logging.
func main() {
	// Connect to the remote SMTP server.
	client, err := smtp.Dial("smtp.smail.com:25")
	if err != nil {
		log.Fatalln(err)
	}
	client.Data()
}
