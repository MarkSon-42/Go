package handler

import (
	"log"
	"net"
)

// ClientHandler로 클라이언트 연결을 처리하는 구조체 선언
type Handler struct {
	Logger *log.Logger
}

func NewHandler(LOGGER *log.Logger) *Handler {
	return &Handler{
		Logger: LOGGER,
	}
}

// 들어오는 클라이언트 연결을 처리하는 메서드
func (h *Handler) HandleConnection(conn net.Conn) {
	defer conn.Close() // 연결 종료 시 conn.Close() 호출

	buffer := make([]byte, 1024)
	for {
		break
		// 임시로 루프 종료 ( 기능 구현중..)
	}

	data := buffer[:n]
}
