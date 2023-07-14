# vmmgo

Implementation of vmm (MemProcFS library)

- only supports windows
- not fully implemented
- tested with 5.7 version

implemented functions are:
- basic functions like initialize, close, free, ...
- modules / pid list
- fully scatter support
- Memory read and write
- Get Vad

note: before use, remember download the binaries of MemProcFS (https://github.com/ufrisk/MemProcFS/releases) and run the example with the full path of vmm.dll as command line argument

### example

```go
// list all modules of windows pid 4 (System)
args := []string{"", "-device", "fpga"}
vmm, _ := vmmgo.Initialize(args)
modules, _ := vmm.GetAllModules(4)
for _, moduleentry := range modules.PMap {
    fmt.Println(moduleentry, *moduleentry.WszText)
}
```

Thanks to Ulf Frisk for MemProcFS (https://github.com/ufrisk/MemProcFS) and Leechcore (https://github.com/ufrisk/LeechCore)