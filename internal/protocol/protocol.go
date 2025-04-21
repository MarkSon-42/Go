package protocol

import (
	"bufio"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

type Message struct {
	Type	string `json:"type"`
	Content string `json:"type"`
)

func HandleJSONProtocol() {
	pass
}


