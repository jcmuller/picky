package main

import (
	"os"
)

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	if len(os.Args) != 2 {
		os.Exit(0)
	}

	arg := os.Args[1]

	c := chooser.New(arg)
	c.Call()
}
