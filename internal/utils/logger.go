// logger.go

package utils

import (
	"log"
	"os"
)

type Logger struct {
	file *os.File
}

func NewLogger(filePath string) (*Logger, error) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	log.SetOutput(file) // log 패키지의 출력 대상을 파일로 설정
	return &Logger{file: file}, nil
}
