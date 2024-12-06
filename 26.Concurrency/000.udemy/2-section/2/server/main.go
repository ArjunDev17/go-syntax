package main

import (
	"io"
	"log"
	"net" //for create server
	"time"
)

func main() {
	//todo :WRITE SERVER PROGRAM TO HANDLE COCURRENT CLIENT CONNECTIONS.
	listner, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatalf("%v :", err)
	}
	for {
		conn, err := listner.Accept()
		if err != nil {
			log.Fatalf("%v :", err)
		}
		go handleConnection(conn) //after this line
	}
}
func handleConnection(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, "response from server \n")
		if err != nil {
			log.Fatalf("%v :", err)
		}
		time.Sleep(time.Second)
	}
}
