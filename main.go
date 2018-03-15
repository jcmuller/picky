package main

import (
	"os"

	"github.com/jcmuller/choosy/chooser"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(0)
	}

	arg := os.Args[1]

	c := chooser.New(arg)
	c.Call()
}
