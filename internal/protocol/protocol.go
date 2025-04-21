package protocol

import (
	"encoding/json"
	"fmt"
	"net"
)

type Message struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

func HandleJSONProtocol(conn net.Conn) error {
	defer conn.Close()

	decoder := json.NewDecoder(conn)
	var msg Message
	if err := decoder.Decode(&msg); err != nil {
		return fmt.Errorf("failed to decode JSON: %v", err)
	}

	fmt.Printf("Received message: %s\n", msg.Content)
	return nil
}
