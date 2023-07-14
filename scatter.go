package vmmgo

import (
	"unsafe"
)

type ScatterInstance struct {
	hS    uintptr
	pid   int32
	flags uintptr
	vmm   *VMM
}

func (inst *VMM) ScatterInitialize(pid int32, flags uintptr) (*ScatterInstance, error) {
	result, _, _ := call("VMMDLL_Scatter_Initialize", initializeResult, uintptr(pid), flags)
	if result == 0 {
		return nil, ERR_CALL
	}

	return &ScatterInstance{hS: result, pid: pid, flags: flags, vmm: inst}, nil
}

func (inst *ScatterInstance) Prepare(va uintptr, size int) error {
	result, _, _ := call("VMMDLL_Scatter_Prepare", inst.hS, va, uintptr(size))
	if result == 0 {
		return ERR_CALL
	}

	return nil
}

func (inst *ScatterInstance) PrepareWrite(va uintptr, data []byte) error {
	result, _, _ := call("VMMDLL_Scatter_PrepareWrite", inst.hS, va, uintptr(unsafe.Pointer(&data[0])), uintptr(len(data)))
	if result == 0 {
		return ERR_CALL
	}

	return nil
}

func (inst *ScatterInstance) Execute() error {
	result, _, _ := call("VMMDLL_Scatter_Execute", inst.hS)
	if result == 0 {
		return ERR_CALL
	}

	return nil
}

func (inst *ScatterInstance) Read(va uintptr, size int) ([]byte, error) {
	pb := make([]byte, size)
	pcbRead := int32(0)
	result, _, _ := call("VMMDLL_Scatter_Read", inst.hS, va, uintptr(size), uintptr(unsafe.Pointer(&pb[0])), uintptr(unsafe.Pointer(&pcbRead)))
	if result == 0 {
		return nil, ERR_CALL
	}

	return pb, nil
}

func (inst *ScatterInstance) Clear() error {
	result, _, _ := call("VMMDLL_Scatter_Clear", inst.hS, uintptr(inst.pid), inst.flags)
	if result == 0 {
		return ERR_CALL
	}

	return nil
}

func (inst *ScatterInstance) CloseHandle() {
	call("VMMDLL_Scatter_CloseHandle", inst.hS)

	inst.hS = 0
	inst.pid = 0
	inst.flags = 0
}
