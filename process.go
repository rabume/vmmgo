package vmmgo

import "unsafe"

func (inst *VMM) ProcessGetAll() (pid []int32, err error) {
	var pcPIDs uintptr
	result, _, _ := call("VMMDLL_PidList", initializeResult, 0, uintptr(unsafe.Pointer(&pcPIDs)))
	if result == 0 {
		// error
		return nil, ERR_CALL
	}

	pids := make([]int32, pcPIDs)
	result, _, _ = call("VMMDLL_PidList", initializeResult, uintptr(unsafe.Pointer(&pids[0])), uintptr(unsafe.Pointer(&pcPIDs)))
	if result == 0 {
		// error
		return nil, ERR_CALL
	}

	return pids, nil
}

func (inst *VMM) PidGetFromName(processName string) (pid int32, err error) {
	szProcName := bytePtrFromString(processName)
	var pdwPID uint32

	result, _, _ := call("VMMDLL_PidGetFromName", initializeResult, uintptr(unsafe.Pointer(szProcName)),
		uintptr(unsafe.Pointer(&pdwPID)))

	if result == 0 {
		return 0, ERR_CALL
	}

	return int32(pdwPID), nil
}
