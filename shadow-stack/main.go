package main

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"runtime/pprof"
	"unsafe"
)

func main() {
	funcValue := reflect.ValueOf(trampoline)
	funcPointer := unsafe.Pointer(funcValue.Pointer())

	call5(func() {
		var i int
		var s shadowStack
		stackwalk(func(f *frame) bool {
			i++
			if i == 4 {
				s.Ret = f.retpc
				f.retpc = uintptr(funcPointer)
				return false
			}
			return true
		})
		setupShadowStack(s)
		fmt.Println("injected")
	})
}

const pprofShadowStackKey = "shadowStack"

func setupShadowStack(s shadowStack) {
	data, err := json.Marshal(s)
	if err != nil {
		panic(err)
	}
	lbls := pprof.Labels(pprofShadowStackKey, string(data))
	ctx := pprof.WithLabels(context.Background(), lbls)
	pprof.SetGoroutineLabels(ctx)
}

func goroutineShadowStack() (s *shadowStack) {
	return extractShadowStack(goroutineLabels())
}

func extractShadowStack(labels map[string]string) (s *shadowStack) {
	if labels == nil {
		return nil
	}
	data, ok := labels[pprofShadowStackKey]
	if !ok {
		return nil
	}

	s = &shadowStack{}
	if err := json.Unmarshal([]byte(data), s); err != nil {
		panic(err)
	}
	return
}

func destroyShadowStack() uintptr {
	labels := goroutineLabels()
	ss := extractShadowStack(labels)
	delete(labels, pprofShadowStackKey)

	var lblStrings []string
	for k, v := range labels {
		lblStrings = append(lblStrings, k, v)
	}
	lbls := pprof.Labels(lblStrings...)
	ctx := pprof.WithLabels(context.Background(), lbls)
	pprof.SetGoroutineLabels(ctx)

	println("destroy shadow stack")

	return ss.Ret
}

type shadowStack struct {
	Ret uintptr
}

func goroutineLabels() map[string]string {
	m := (*map[string]string)(runtime_getProfLabel())
	if m == nil {
		return nil
	}
	return *m
}

//go:linkname runtime_getProfLabel runtime/pprof.runtime_getProfLabel
func runtime_getProfLabel() unsafe.Pointer

func stackwalk(fn func(*frame) bool) {
	frame := (*frame)(unsafe.Pointer(regfp()))
	for {
		if !fn(frame) {
			return
		}
		frame = frame.pointer
		if frame.pointer == nil {
			break
		}
	}
}

type frame struct {
	pointer *frame
	retpc   uintptr
}

// regfp returns the frame pointer addr in the callers frame by
func regfp() uintptr

func trampoline()

func callN(depth int, fn func()) {
	defer func() { fmt.Println("ret", depth) }()
	if depth > 0 {
		callN(depth-1, fn)
	} else {
		fn()
	}
}

//go:noinline
func call5(fn func()) { call4(fn); println("ret 5") }

//go:noinline
func call4(fn func()) { call3(fn); println("ret 4") }

//go:noinline
func call3(fn func()) { call2(fn); println("ret 3") }

//go:noinline
func call2(fn func()) { call1(fn); println("ret 2") }

//go:noinline
func call1(fn func()) { call0(fn); println("ret 1") }

//go:noinline
func call0(fn func()) { fn(); println("ret 0") }
