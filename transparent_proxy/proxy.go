package main

import (
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen on port 8080: %v", err)
	}
	defer listener.Close()

	log.Println("Listening on port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept connection: %v", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(src net.Conn) {
	defer src.Close()

	// should parse the destination from the incoming request in real implementation
	destAddr := "localhost:3000"

	dest, err := net.Dial("tcp", destAddr)
	if err != nil {
		log.Printf("Failed to connect to destination %s: %v", destAddr, err)
		return
	}
	defer dest.Close()

	log.Printf("Proxying from %s to %s", src.RemoteAddr(), destAddr)

	// here we can do any kind of processing on the data

	// forward data in both directions
	go io.Copy(dest, src)
	io.Copy(src, dest)
}
