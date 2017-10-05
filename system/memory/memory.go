package memory

import "runtime"

// Returns Bytes
func MemAllocated() uint64 {
	runtime.GC()
	m := new(runtime.MemStats)
	runtime.ReadMemStats(m)
	return m.HeapSys
}

// Returns Bytes
func MemInUse() uint64 {
	runtime.GC()
	m := new(runtime.MemStats)
	runtime.ReadMemStats(m)
	return m.HeapInuse
}
