package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	//todo :Connect to the server on localhost port 8000
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatalf("%v :", err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)

}
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
