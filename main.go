// Package main is the entrypoint
package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/jcmuller/picky/chooser"
	"github.com/jcmuller/picky/config"
	"github.com/jcmuller/picky/configfile"
)

func onFileError() {
	errorString := "http://juancmuller.com/simplemessage/pickyerror.html?home=%s"
	err := exec.Command("chromium-browser", fmt.Sprintf(errorString, os.Getenv("HOME"))).Run()
	if err != nil {
		panic(err)
	}
}

func main() {
	if len(os.Args) < 2 {
		os.Exit(0)
	}

	arg := strings.Join(os.Args[1:], " ")

	configFilePath, err := configfile.FilePath(os.Getenv("HOME"))
	handle(err)

	configYaml, err := configfile.FileContents(configFilePath, onFileError)

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
