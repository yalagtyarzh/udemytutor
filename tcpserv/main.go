package main

import (
	"io"
	"log"
	"net"
)

func handle(conn net.Conn) {
	defer conn.Close()
	io.WriteString(conn, "I see you conntected")
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		go handle(conn)
	}
}
