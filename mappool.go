package vmmgo

import (
	"unsafe"
)

func (inst *VMM) MapPool(flag uintptr) (*VMMDLL_MAP_POOL, error) {
	var ptr uintptr
	result, _, _ := call("VMMDLL_Map_GetPool", initializeResult, uintptr(unsafe.Pointer(&ptr)), flag)
	if result == 0 {
		return nil, ERR_CALL
	}

	oneElement := (*VMMDLL_MAP_POOL_oneelement)(unsafe.Pointer(ptr))

	if oneElement.DwVersion != VMMDLL_MAP_POOL_VERSION {
		return nil, ERR_BAD_VERSION
	}

	mapPool := &VMMDLL_MAP_POOL{
		PMap:       make([]VMMDLL_MAP_POOLENTRY, oneElement.CMap),
		CMap:       oneElement.CMap,
		_Reserved1: oneElement._Reserved1,
		CbTotal:    oneElement.CbTotal,
		CTag:       oneElement.CTag,
		DwVersion:  oneElement.DwVersion,
		PiTag2Map:  oneElement.PiTag2Map,
		PTag:       oneElement.PTag,
	}

	sizeofchild := unsafe.Sizeof(VMMDLL_MAP_POOLENTRY{})

	address := uintptr(unsafe.Pointer(&oneElement.PMap[0]))
	for i := 0; i < int(oneElement.CMap); i++ {
		t := (*VMMDLL_MAP_POOLENTRY)(unsafe.Pointer(address + (uintptr(i) * sizeofchild)))

		mapPool.PMap[i] = *t
	}

	call("VMMDLL_MemFree", uintptr(ptr))

	return mapPool, nil
}
