package bumpy

type MyStruct struct {
	data [8]byte
}

func NewAllocator[T any](size int) *Allocator[T] {
	return &Allocator[T]{size: size}
}

type Allocator[T any] struct {
	list []T
	size int
	i    int
}

func (t *Allocator[T]) Alloc() *T {
	if len(t.list) == 0 || t.i == t.size {
		t.list = make([]T, t.size)
		t.i = 0
	}
	a := &t.list[t.i]
	t.i++
	return a
}

func NewAllocatorMyStruct(size int) *AllocatorMyStruct {
	return &AllocatorMyStruct{size: size}
}

type AllocatorMyStruct struct {
	list []MyStruct
	size int
	i    int
}

func (t *AllocatorMyStruct) Alloc() *MyStruct {
	if len(t.list) == 0 || t.i == t.size {
		t.list = make([]MyStruct, t.size)
		t.i = 0
	}
	a := &t.list[t.i]
	t.i++
	return a
}
