package vmmgo

import (
	"errors"
	"log"
)

var (
	initializeResult uintptr
)

func call(functionName string, a ...uintptr) (r1, r2 uintptr, lastErr error) {
	log.Println("LINUX NOT IMPLEMENTED YET")
	return 0, 0, errors.New("NOT IMPLEMENTED YET")
}
