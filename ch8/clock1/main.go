// Clock1 is a TCP server that periodically displays the time.
package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8005")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // For example, connection failure
			continue
		}
		handleConn(conn) // single connection processing
	}
}

func handleConn(c net.Conn) {
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // for example, disconnecting the client
		}
		time.Sleep(1 * time.Second)
	}
}
