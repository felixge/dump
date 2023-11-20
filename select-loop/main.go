package main

import (
	"os"
	"runtime/pprof"
	"time"
)

func main() {
	file, _ := os.Create("./cpu.pprof")
	pprof.StartCPUProfile(file)
	defer pprof.StopCPUProfile()

	go func() {
ch1 := make(chan struct{})
ch2 := make(chan struct{})
for {
	select {
	case <-ch1:
	case <-ch2:
	case <-time.After(10 * time.Millisecond):
		// intentionally blank
	}
}
	}()

	time.Sleep(3 * time.Second)
}
