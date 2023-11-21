package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"time"
	"unsafe"
)

type Node struct {
	Next *Node
	Val  int
}

func main() {
	kB := 1
	maxkB := 1024 * 1024 * kB
	w := csv.NewWriter(os.Stdout)
	defer w.Flush()
	w.Write([]string{"KiB", "nsPerOp", "GiB/s"})

	iter := 10
	s := unsafe.Sizeof(Node{})
	for kB <= maxkB {
		n := (1024 * kB) / int(s)
		nsPerOp := bench(n, iter)
		w.Write([]string{
			fmt.Sprintf("%d", kB),
			fmt.Sprintf("%f", nsPerOp),
			fmt.Sprintf("%f", 1e9/nsPerOp*float64(s)/1024/1024/1024),
		})
		w.Flush()
		kB *= 2
	}
}

func bench(n, iters int) float64 {
	head := makeList(n)
	runtime.GC()
	defer runtime.GC()
	start := time.Now()
	for i := 0; i < iters; i++ {
		for node := head; node != nil; node = node.Next {
		}
	}
	return float64(time.Since(start)) / float64(n*iters)
}

// makeList makes a linked list of length n that is laid out in memory in a
// random order. It returns a pointer to the first node.
func makeList(n int) *Node {
	nodes := make([]Node, n)
	for i := range nodes {
		nodes[i].Val = i
	}
	pointers := make([]*Node, len(nodes))
	for i := range pointers {
		pointers[i] = &nodes[i]
	}
	Shuffle(pointers)

	for i := 0; i < len(pointers)-1; i++ {
		pointers[i].Next = pointers[i+1]
	}
	pointers[len(pointers)-1].Next = nil
	return pointers[0]
}

// Shuffle shuffles the given slice using the Fisher-Yates algorithm.
func Shuffle[T any](slice []T) {
	for i := len(slice) - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}
