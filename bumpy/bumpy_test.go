package bumpy

import "testing"

var sink *MyStruct

func BenchmarkAlloc(b *testing.B) {
	b.Run("bumpy-generic", func(b *testing.B) {
		a := NewAllocator[MyStruct](100)
		for i := 0; i < b.N; i++ {
			sink = a.Alloc()
		}
	})

	b.Run("bumpy-direct", func(b *testing.B) {
		a := NewAllocatorMyStruct(100)
		for i := 0; i < b.N; i++ {
			sink = a.Alloc()
		}
	})

	b.Run("go", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sink = &MyStruct{}
		}
	})
}
