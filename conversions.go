package vmmgo

import (
	"syscall"
	"unsafe"
)

func stringSlicePtr(ss []string) []*byte {
	bb := make([]*byte, len(ss))
	for i := 0; i < len(ss); i++ {
		bb[i], _ = syscall.BytePtrFromString(ss[i])
	}
	return bb
}

func bytePtrFromString(input string) *byte {
	result, _ := syscall.BytePtrFromString(input)
	return result
}

func byteSliceFromString(s string) ([]byte, error) {
	for i := 0; i < len(s); i++ {
		if s[i] == 0 {
			return nil, syscall.EINVAL
		}
	}
	a := make([]byte, len(s)+1)
	copy(a, s)
	return a, nil
}

func readString(pointer unsafe.Pointer) *string {
	myPointer := pointer
	result := make([]byte, 0)
	for {
		v := *((*byte)(myPointer))
		if v == 0 {
			break
		}

		myPointer = unsafe.Pointer(uintptr(myPointer) + 1)
		result = append(result, v)
	}
	rStr := string(result)
	return &rStr
}
