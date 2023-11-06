package memStatus

import (
	"fmt"
	"runtime"
	"time"
)

func Memstatus() {
	start := time.Now()
	var m runtime.MemStats
	elapsed := time.Since(start)
	runtime.ReadMemStats(&m)
	memory := m.Alloc / 1024 / 1024
	if memory == 0 {
		memory = m.Alloc / 1024
		fmt.Printf("Memory used: %v kb \n", memory)
	} else {
		fmt.Printf("Memory used: %v MiB \n", memory)
	}
	fmt.Printf("function took %s \n", elapsed)
}
