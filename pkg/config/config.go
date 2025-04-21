package config

import (
	"os"
)

// server 설정 저장하는 구조체 Config
type Config struct {
	Address string
}

func LoadConfig() Config {
	address := os.Getenv("SERVER_ADDRESS")
	if address == "" {
		address = "localhost:9000"
	}

	return Config{
		Address: address,
	}
}

