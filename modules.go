package vmmgo

import (
	"unsafe"
)

func (inst *VMM) GetAllModules(pid int32) (*VMMDLL_MAP_MODULE, error) {
	pcbMap := uintptr(0)
	result, _, _ := call("VMMDLL_Map_GetModuleU", initializeResult, uintptr(pid), uintptr(unsafe.Pointer(&pcbMap)))
	if result == 0 {
		// error
		return nil, ERR_CALL
	}

	v4 := (*vMMDLL_MAP_MODULE)(unsafe.Pointer(pcbMap))

	v5 := &VMMDLL_MAP_MODULE{
		PMap:        make([]VMMDLL_MAP_MODULEENTRY, v4.CMap),
		CMap:        v4.CMap,
		_Reserved1:  v4._Reserved1,
		CbMultiText: v4.CbMultiText,
		PbMultiText: v4.PbMultiText,
		DwVersion:   v4.DwVersion,
	}

	sizeofchild := unsafe.Sizeof(VMMDLL_MAP_MODULEENTRY{})
	add := uintptr(unsafe.Pointer(&v4.PMap))
	for i := 0; i < int(v4.CMap); i++ {
		t := (*VMMDLL_MAP_MODULEENTRY)(unsafe.Pointer(add + (uintptr(i) * sizeofchild)))
		t.WszText = readString(unsafe.Pointer(t.WszText))
		t.WszFullName = readString(unsafe.Pointer(t.WszFullName))

		v5.PMap[i] = *t
	}

	return v5, nil
}
