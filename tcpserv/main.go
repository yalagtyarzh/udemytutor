package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func handle(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	io.WriteString(conn, "I see you connected\n")
	for scanner.Scan() {
		ln := scanner.Text()
		if ln == "" {
			break
		}

		fmt.Println(ln)
		io.WriteString(conn, fmt.Sprint(ln, "\n"))
	}

	fmt.Println("Code got here")
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
