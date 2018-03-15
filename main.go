package main

import (
	"os"
	"strings"

	"github.com/jcmuller/choosy/chooser"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(0)
	}

	arg := strings.Join(os.Args[1:], " ")

	c := chooser.New(arg)
	c.Call()
}
