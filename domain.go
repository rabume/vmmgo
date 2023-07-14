package vmmgo

type vMMDLL_MAP_MODULE struct { // total size is 40 (0x28)
	DwVersion   uint32
	_Reserved1  [5]uint32
	PbMultiText uintptr
	CbMultiText uint32
	CMap        uint32
	PMap        [1024]VMMDLL_MAP_MODULEENTRY
}

type VMMDLL_MAP_MODULE struct { // total size is 40 (0x28)
	DwVersion   uint32
	_Reserved1  [5]uint32
	PbMultiText uintptr
	CbMultiText uint32
	CMap        uint32
	PMap        []VMMDLL_MAP_MODULEENTRY
}

type VMMDLL_MAP_MODULEENTRY struct {
	VaBase         uint64
	VaEntry        uint64
	CbImageSize    uint32
	FWoW64         uint32
	WszText        *string
	_Reserved3     uint32
	_Reserved4     uint32
	WszFullName    *string
	Tp             uint32
	CbFileSizeRaw  uint32
	CSection       uint32
	CEAT           uint32
	CIAT           uint32
	_Reserved2     uint32
	_Reserved1     [3]uint64
	pExDebugInfo   uintptr
	pExVersionInfo uintptr
}

type VMMDLL_MAP_VAD struct {
	DwVersion   uint32
	_Reserved1  [4]uint32
	CPage       uint32
	PbMultiText uintptr
	CbMultiText uint32
	CMap        uint32
	PMap        []VMMDLL_MAP_VADENTRY
}

type vMMDLL_MAP_VAD struct {
	DwVersion   uint32
	_Reserved1  [4]uint32
	CPage       uint32
	PbMultiText uintptr
	CbMultiText uint32
	CMap        uint32
	PMap        [1024]VMMDLL_MAP_VADENTRY
}

type VMMDLL_MAP_VADENTRY struct {
	VaStart         uintptr
	VaEnd           uintptr
	VaVad           uintptr
	Dw0             uint32
	Dw1             uint32
	Dwu2            uint32
	CbPrototypePte  uint32
	VaPrototypePte  uintptr
	VaSubsection    uintptr
	UszText         *string
	_FutureUse1     uint32
	_Reserved1      uint32
	VaFileObject    uintptr
	CVadExPages     uint32
	CVadExPagesBase uint32
	_Reserved2      uintptr
}

type VMMDLL_MAP_POOL struct {
	DwVersion  uint32
	_Reserved1 [6]uint32
	CbTotal    uint32
	PiTag2Map  *uint32
	PTag       *VMMDLL_MAP_POOLENTRYTAG
	CTag       uint32
	CMap       uint32
	PMap       []VMMDLL_MAP_POOLENTRY
}

type VMMDLL_MAP_POOL_oneelement struct {
	DwVersion  uint32
	_Reserved1 [6]uint32
	CbTotal    uint32
	PiTag2Map  *uint32
	PTag       *VMMDLL_MAP_POOLENTRYTAG
	CTag       uint32
	CMap       uint32
	PMap       [1024]VMMDLL_MAP_POOLENTRY
}

type VMMDLL_MAP_POOLENTRYTAG struct {
	// size: 16
	DwTag    uint32
	_Filler  uint32
	CEntry   uint32
	ITag2Map uint32
}

type VMMDLL_MAP_POOLENTRY struct {
	// total size: 24
	VA            uintptr
	DwTag         uint32
	_ReservedZero byte
	FAlloc        byte
	TpPool        byte
	TpSS          byte
	Cb            uint32
	_Filler       uint32
}
