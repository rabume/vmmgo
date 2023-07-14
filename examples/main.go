package main

import (
	"fmt"
	"github.com/pineda89/vmmgo"
	"strings"
)

const target_process = "notepad.exe"

func main() {
	args := []string{"", "-device", "fpga"}
	vmm, err := vmmgo.Initialize(args)
	checkErr(err)

	pids, err := vmm.ProcessGetAll()
	checkErr(err)
	fmt.Println("pids", pids)

	pid, err := vmm.PidGetFromName(target_process)
	checkErr(err)
	fmt.Println("process pid", pid)

	modules, err := vmm.GetAllModules(pid)
	checkErr(err)
	var baseAddress uintptr

	for i := range modules.PMap {
		if strings.Contains(*modules.PMap[i].WszText, target_process) {
			baseAddress = uintptr(modules.PMap[i].VaBase)
		}
		fmt.Println(i, modules.PMap[i], *modules.PMap[i].WszText)
	}
	fmt.Println("base address is", baseAddress)

	data, err := vmm.MemReadEx(pid, baseAddress, 4, vmmgo.FLAG_NOCACHE|vmmgo.FLAG_NOCACHEPUT|vmmgo.FLAG_NOPAGING|vmmgo.FLAG_NOPAGING_IO)
	checkErr(err)
	fmt.Println("baseAddress data is", data)

	fmt.Println("writting in base address...")
	err = vmm.MemWrite(pid, baseAddress, []byte{0, 1, 2, 3})
	checkErr(err)

	data, err = vmm.MemReadEx(pid, baseAddress, 4, vmmgo.FLAG_NOCACHE|vmmgo.FLAG_NOCACHEPUT|vmmgo.FLAG_NOPAGING|vmmgo.FLAG_NOPAGING_IO)
	checkErr(err)
	fmt.Println("baseAddress data after write is", data)

	mapPool, err := vmm.MapPool(vmmgo.VMMDLL_POOLMAP_FLAG_ALL)
	checkErr(err)
	for i := range mapPool.PMap {
		fmt.Println(i, mapPool.PMap[i].VA, mapPool.PMap[i].Cb)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
