// Package main is the entrypoint
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jcmuller/choosy/chooser"
	"github.com/jcmuller/choosy/config"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(0)
	}

	arg := strings.Join(os.Args[1:], " ")

	configFilePath, err := config.FilePath()
	handle(err)

	configYaml, err := config.FileContents(configFilePath)

	if err != nil {
		fmt.Println(fmt.Errorf("%s", err))
		os.Exit(1)
	}

	config, err := config.New(configYaml)
	handle(err)

	c := chooser.New(config, arg)
	c.Call()
}

func handle(err error) {
	if err != nil {
		panic(err)
	}
}
