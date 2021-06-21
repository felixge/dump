package main

import (
	"fmt"
	"os"

	plugalt "github.com/felixge/dump/go-plugin-alt"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	fmt.Printf("Registered Plugins:\n")
	for _, p := range plugalt.Plugins() {
		fmt.Printf("%#v\n", p)
	}
	return nil
}

type Plugin struct {
	Name string
}

func RegisterPlugin(p Plugin) {
}

var plugins []Plugin
