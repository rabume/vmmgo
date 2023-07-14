package vmmgo

import (
	"log"
	"sync"
	"unsafe"
)

type VMM struct {
}

var mux sync.RWMutex

func Initialize(arguments []string) (*VMM, error) {
	argc := uint32(len(arguments))
	argv := stringSlicePtr(arguments)

	result, _, _ := call("VMMDLL_Initialize", uintptr(argc), uintptr(unsafe.Pointer(&argv[0])))
	if result == 0 {
		return nil, ERR_CALL
	}
	log.Println("initialize ok")
	initializeResult = result

	result, _, _ = call("VMMDLL_InitializePlugins", initializeResult)
	if result == 0 {
		return nil, ERR_CALL
	}
	log.Println("initialize plugins ok")

	return &VMM{}, nil
}

func (inst *VMM) Close() bool {
	result, _, _ := call("VMMDLL_Close", initializeResult)
	return result == 0
}
