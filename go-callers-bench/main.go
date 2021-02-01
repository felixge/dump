package main

import (
	"flag"
	"fmt"
	"runtime"
	"time"
)

func main() {
	var (
		ops   = flag.Int("ops", 100000, "Number of times to call runtime.Callers().")
		depth = flag.Int("depth", 32, "Stack depth at which to call runtime.Callers()")
	)
	flag.Parse()

	dt := bench(*depth, *ops)
	fmt.Printf("%s/op\n", dt)
}

// bench returns the avg time/op for calling runtime.Callers() ops times at the
// given stack depth.
func bench(depth, ops int) time.Duration {
	pcs := make([]uintptr, depth*10)
	start := time.Now()
	for i := 0; i < ops; i++ {
		n := runtime.Callers(1, pcs)
		if n > depth {
			panic("bad")
		} else if n < depth {
			return bench(depth, ops)
		}
	}
	return time.Since(start) / time.Duration(ops)
}

func printPcs(pcs []uintptr) {
	frames := runtime.CallersFrames(pcs)
	for {
		frame, more := frames.Next()
		fmt.Printf("%s\n", frame.Function)
		if !more {
			break
		}
	}
}
