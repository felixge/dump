package main

import (
	"flag"
	"fmt"
	"time"
)

func cputicks() int64

func main() {
	var (
		ops = flag.Int64("ops", 10000000, "Number of times to call cputicks.")
	)
	flag.Parse()
	start := time.Now()
	for i := int64(0); i < *ops; i++ {
		cputicks()
	}
	dt := time.Since(start)
	fmt.Printf("%d ops in %s\n", *ops, dt)
	fmt.Printf("%.1f ns/op\n", float64(dt.Nanoseconds())/float64(*ops))
}
