/*
$ go version
go version go1.15.6 linux/amd64
$ go build hello_world.go
$ ldd hello_world
	not a dynamic executable
*/

package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}
