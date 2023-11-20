package main

import (
	"testing"
)

type S struct {
	a, b, c int64
	d, e, f string
	g, h, i float64
}

func (s S) stack(s1 S) {}

func (s *S) heap(s1 *S) {}

func byCopy() S {
	return S{
		a: 1, b: 1, c: 1,
		e: "foo", f: "foo",
		g: 1.0, h: 1.0, i: 1.0,
	}
}

func byPointer() *S {
	return &S{
		a: 1, b: 1, c: 1,
		e: "foo", f: "foo",
		g: 1.0, h: 1.0, i: 1.0,
	}
}

func BenchmarkMemoryHeap(b *testing.B) {
	var s *S
	var s1 *S

	s = byPointer()
	s1 = byPointer()
	for i := 0; i < b.N; i++ {
		for i := 0; i < 1000000; i++ {
			s.heap(s1)
		}
	}
}

func BenchmarkMemoryStack(b *testing.B) {
	var s S
	var s1 S

	s = byCopy()
	s1 = byCopy()
	for i := 0; i < b.N; i++ {
		for i := 0; i < 1000000; i++ {
			s.stack(s1)
		}
	}
}
