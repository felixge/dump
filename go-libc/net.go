/*
$ go version
go version go1.15.6 linux/amd64
$ go build net.go
$ ldd net
	linux-vdso.so.1 (0x00007fff9a7cf000)
	libpthread.so.0 => /lib/x86_64-linux-gnu/libpthread.so.0 (0x00007fea52e49000)
	libc.so.6 => /lib/x86_64-linux-gnu/libc.so.6 (0x00007fea52c88000)
	/lib64/ld-linux-x86-64.so.2 (0x00007fea52e70000)
*/

package main

import (
	"fmt"
	"net"
)

func main() {
	var a net.Addr
	fmt.Println("Hello World", a)
}
