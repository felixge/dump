// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

/*
void burn_cgo() ;
static void *burn_pthread(void *arg);
static void *burn_pthread_go(void *arg);
void prep_cgo();
void prep_pthread();
void prep_pthread_go();
*/
import "C"

import (
	"flag"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

func main() {
	goThread := flag.Int("go-threads", 0, "Number of Go-created threads to run Go code")
	cgoThread := flag.Int("cgo-threads", 0, "Number of Go-created threads to run cgo code")
	cgoPthread := flag.Int("cgo-pthreads", 0, "Number of cgo-created threads to run cgo code")
	goPthread := flag.Int("go-pthreads", 0, "Number of cgo-created threads to run Go code")
	sleep := flag.Duration("sleep", 2*time.Second, "Duration of test")
	presleep := flag.Duration("presleep", 0, "How long to wait for threads to start before profiling begins")
	cpuprofile := flag.String("cpuprofile", "", "Write a CPU profile to the specified file")
	flag.Parse()

	startProfile := func() {}
	if name := *cpuprofile; name != "" {
		f, err := os.Create(name)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		startProfile = func() {
			err := pprof.StartCPUProfile(f)
			if err != nil {
				log.Fatal(err)
			}
		}
		defer pprof.StopCPUProfile()
	}

	if *presleep <= 0 {
		startProfile()
	}

	for i := 0; i < *goThread; i++ {
		go in_go()
	}
	for i := 0; i < *cgoThread; i++ {
		go in_cgo()
	}
	for i := 0; i < *cgoPthread; i++ {
		go in_pthread()
	}
	for i := 0; i < *goPthread; i++ {
		go in_pthread_go()
	}

	if *presleep > 0 {
		time.Sleep(*presleep)
		startProfile()
	}

	time.Sleep(*sleep)
}

func in_go() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	for {
	}
}

func in_cgo() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	C.prep_cgo()
}

func in_pthread() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	C.prep_pthread()
}

func in_pthread_go() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	C.prep_pthread_go()
}

//export burn_callback
func burn_callback() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	for {
	}
}