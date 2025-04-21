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
	log.SetOutput(file)                 // log 패키지의 출력 대상을 파일로 설정
	log.SetFlags(log.Ldate | log.Ltime) // 로그에 날짜, 시간, 파일명 포함.
	log.Println("Logger initialized.")  // 로그 초기화 메시지
	return &Logger{file: file}, nil
}

func (l *Logger) Close() {
	if l.file != nil {
		l.file.Close() // 파일 닫기
	}
}

func (l *Logger) Info(msg string) {
	log.Println("INFO:", msg)
}

func (l *Logger) Error(msg string) {
	log.Println("ERROR:", msg)
}
