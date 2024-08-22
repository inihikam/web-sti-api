package utils

import (
	"log"
)

func LogInfo(msg string) {
	log.Printf("[INFO] %s", msg)
}

func LogError(err error) {
	log.Printf("[ERROR] %s", err.Error())
}