// Clock2 is a TCP server that periodically displays the time with parallel connection processing.
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter port number: ")
	// port, _ := reader.ReadString('\n')
	var port int
	fmt.Print("Enter port number: ")
	_, err := fmt.Scanf("%d", &port)
	if err != nil {
		log.Print(err)
	}

	listener, err := net.Listen("tcp", "localhost:"+strconv.Itoa(port))
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // For example, connection failure
			continue
		}
		go handleConn(conn) // parallel connection processing
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
