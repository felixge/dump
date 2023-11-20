package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	var mu sync.Mutex
	var count int

	go func() {
		count++
		mu.Lock()
		mu.Unlock()
	}()
	go func() {
		time.Sleep(time.Millisecond)
		mu.Lock()
		mu.Unlock()
		count++
	}()

	time.Sleep(200 * time.Millisecond)
	return nil
}

var foo int
var fooMu sync.Mutex

func run2() error {
	a := NewVectorClock()
	b := NewVectorClock()
	c := NewVectorClock()
	//initialA := a
	a = a.Send("A")
	b = b.Send("B")
	//b = b.Receive("B", a)
	//b = b.Send("B")
	//c = c.Receive("C", b)
	fmt.Printf("a.clocks: %v\n", a.clocks)
	fmt.Printf("b.clocks: %v\n", b.clocks)
	fmt.Printf("c.clocks: %v\n", c.clocks)

	fmt.Printf("a.Before(b): %v\n", a.Before(b))
	fmt.Printf("a.Before(b): %v\n", a.Before(b))
	fmt.Printf("a.Before(b): %v\n", b.Before(a))

	//return nil

	//var wg sync.WaitGroup
	//wg.Add(2)
	//go func() {
	//defer wg.Done()
	//g1()
	//}()
	//go func() {
	//defer wg.Done()
	//g2()
	//}()
	//wg.Wait()
	//return nil
	return nil
}

//func g1() {
//runtime.RaceAcquire(unsafe.Pointer(&foo))
//runtime.RaceRelease(unsafe.Pointer(&foo))
//runtime.RaceRead(unsafe.Pointer(&foo))

////time.Sleep(time.Second)
////fooMu.Lock()
////defer fooMu.Unlock()
////foo = 1
//}

//func g2() {
//runtime.RaceAcquire(unsafe.Pointer(&foo))
//runtime.RaceWrite(unsafe.Pointer(&foo))
////runtime.RaceRelease(unsafe.Pointer(&foo))

//time.Sleep(time.Second)
////runtime.RaceRelease(unsafe.Pointer(&foo))

////time.Sleep(time.Second)
////fooMu.Lock()
////defer fooMu.Unlock()
////foo = 2
//}

func NewVectorClock() *VectorClock {
	return &VectorClock{clocks: map[string]int{}}
}

type VectorClock struct {
	clocks map[string]int
}

func (vc *VectorClock) Send(process string) *VectorClock {
	clone := vc.clone()
	clone.clocks[process]++
	return clone
}

func (vc *VectorClock) Receive(process string, fromClock *VectorClock) *VectorClock {
	clone := vc.Send(process)
	for k, v := range fromClock.clocks {
		if v > vc.clocks[k] {
			clone.clocks[k] = v
		}
	}
	return clone
}

func (vc *VectorClock) Before(other *VectorClock) bool {
	oneStrictLt := false
	for k, v := range vc.clocks {
		if v > other.clocks[k] {
			return false
		} else if v < other.clocks[k] {
			oneStrictLt = true
		}
	}
	for k, v := range other.clocks {
		if v > vc.clocks[k] {
			oneStrictLt = true
		}
	}
	return oneStrictLt
}

func (vc *VectorClock) Concurrent(other *VectorClock) bool {
	return !(vc.Before(other) || other.Before(vc))
}

func (vc *VectorClock) clone() *VectorClock {
	clone := NewVectorClock()
	for k, v := range vc.clocks {
		clone.clocks[k] = v
	}
	return clone
}
