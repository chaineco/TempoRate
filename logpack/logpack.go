package logpack

import (
	"log"
)

// Simple Info Msg
func Info(msg string, funcName string) {
	log.Println("INFO:", funcName, "-", msg)
}

// Simple Info Msg
func Error(msg string, funcName string) {
	log.Println("ERROR:", funcName, "-", msg)
}

// Simple Info Msg
func Warn(msg string, funcName string) {
	log.Println("WARNING:", funcName, "-", msg)
}

// Simple Info Msg
func Fatal(msg string, funcName string) {
	log.Println("FATAL:", funcName, "-", msg)
}
