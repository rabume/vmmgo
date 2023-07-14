package vmmgo

import "unsafe"

func (inst *VMM) MemReadEx(pid int32, va uintptr, size uint32, flags uint32) ([]byte, error) {
	output := make([]byte, size)
	outputLen := uintptr(0)
	result, _, _ := call("VMMDLL_MemReadEx", initializeResult, uintptr(pid), va, uintptr(unsafe.Pointer(&output[0])), uintptr(size), uintptr(unsafe.Pointer(&outputLen)), uintptr(flags))

	if result == 0 {
		return []byte{}, ERR_CALL
	}

	return output, nil
}

func (inst *VMM) MemReadExPointer(pid int32, va uintptr, size uint32, flags uint32) (uintptr, error) {
	var output uintptr
	outputLen := uintptr(0)
	result, _, _ := call("VMMDLL_MemReadEx", initializeResult, uintptr(pid), va, uintptr(unsafe.Pointer(&output)), uintptr(size), uintptr(unsafe.Pointer(&outputLen)), uintptr(flags))

	if result == 0 {
		return output, ERR_CALL
	}

	return output, nil
}
