// cmd/server/main.go
package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	address := "localhost:9000"
	fmt.Println("starting server..:", address)

	// TCP 리스너 생성
	listener, err := net.Listen("TCP", address)
	if err != nil {
		log.Fatalf("failed server staring.: %v", err)
	}
	defer listener.Close()

	fmt.Println("wait client connection.")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("failed client connection.: %v", err)
			continue
		}
		fmt.Println("client connected.:", conn.RemoteAddr())
		conn.Close()
	}
}
