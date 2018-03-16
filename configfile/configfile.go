package configfile

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

var configFileTemplate = "%s/.config/picky/config"

// FilePath returns the config file path
func FilePath(home string) (file string, err error) {
	file = fmt.Sprintf(configFileTemplate, home)
	return
}

type fn func()

// OnFileError handles error
func OnFileError() {
	errorString := "http://juancmuller.com/simplemessage/pickyerror.html?home=%s"
	err := exec.Command("chromium-browser", fmt.Sprintf(errorString, os.Getenv("HOME"))).Run()
	if err != nil {
		panic(err)
	}
}

// FileContents reads the config file
func FileContents(path string, onFileError fn) (configFile []byte, err error) {
	configFile, err = ioutil.ReadFile(path)
	if os.IsNotExist(err) {
		onFileError()
		return nil, errors.New("config not found")
	}

	return
}
