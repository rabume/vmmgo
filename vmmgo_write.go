package vmmgo

import "unsafe"

func (inst *VMM) MemWrite(pid int32, va uintptr, data []byte) error {
	result, _, _ := call("VMMDLL_MemWrite", initializeResult, uintptr(pid), va, uintptr(unsafe.Pointer(&data[0])), uintptr(len(data)))

	if result == 0 {
		return ERR_CALL
	}

	return nil
}
