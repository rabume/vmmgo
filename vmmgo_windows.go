package vmmgo

import (
	"os"
	"syscall"
)

var (
	vmmdll           = syscall.NewLazyDLL(findDllPath())
	functions        = make(map[string]*syscall.LazyProc)
	initializeResult uintptr
)

func findDllPath() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	}
	return "vmm.dll"
}

func getFunction(functionName string) *syscall.LazyProc {
	mux.Lock()
	defer mux.Unlock()

	if f, ok := functions[functionName]; !ok {
		f = vmmdll.NewProc(functionName)
		functions[functionName] = f
		return f
	} else {
		return f
	}
}

func call(functionName string, a ...uintptr) (r1, r2 uintptr, lastErr error) {
	return getFunction(functionName).Call(a...)
}
