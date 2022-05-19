package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

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

func handle(conn net.Conn) {
	defer conn.Close()
	request(conn)
}

func request(conn net.Conn) {
	body := "xd"
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/plain\r\n")
	io.WriteString(conn, "\r\n")
}
