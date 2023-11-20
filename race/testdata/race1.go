package main

import (
	"fmt"
	"sync"
)

var x = 0
var mu sync.Mutex

func main() {
	go func() {
		x = 5
		mu.Lock()
		mu.Unlock()
	}()
	go func() {
		mu.Lock()
		mu.Unlock()
		_ = x
	}()
}

func main3() {
	go func() { x = 5 }()
	go func() { _ = x }()
}

func dataRace() {
	go func() { x = 1 }()
	go func() { fmt.Println(x) }()
}

// DRAndRC contains a data race and a race condition.
func DR_RC() {
	x := 0
	go func() { x = 1 }()
	fmt.Println(x)
}

// RC contains a race condition, but no data race.
func RC() {
	var mu sync.Mutex
	x := 0
	go func() {
		mu.Lock()
		x = 1
		mu.Unlock()
	}()
	mu.Lock()
	fmt.Println(x)
	mu.Unlock()
}

func dataRaceButNotRaceCondition() {
	x := 0
	go func() { x = 0 }()
	fmt.Println(x)
}
