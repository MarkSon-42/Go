// cmd/server/main.go
package main

import (
	"fmt" // go에서 쓰는 std in,out 패키지
	"log" // log 패키지
	"net" // TCP/IP 통신을 위한 패키지
)

func main() {
	address := "localhost:9000"
	fmt.Println("starting server..:", address)

	// TCP 리스너 생성
	listener, err := net.Listen("TCP", address)
	if err != nil {
		log.Fatalf("failed server staring.: %v", err) // 리스너 생성 실패하면 프로그램 종료
	}
	defer listener.Close() // main 함수 종료시 리스너 닫음.
	// defer? : defer는 함수가 종료될 때까지 실행되지 않음.

	// defer -> 지연 실행을 위해 사용되는 키워드. 특정 함수 호출을 현재 함수가 종료될 때 실행하도록 예약
	// >> 파일 닫기나... DB 연결 종료와 같은 작업
	// stack구조로 활용가능 LIFO
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
