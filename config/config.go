package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/jcmuller/choosy/browser"
	"github.com/jcmuller/choosy/rule"

	yaml "gopkg.in/yaml.v2"
)

// Config is a struct that holds configuration
type Config struct {
	Default  *rule.Rule                  `yaml:"default"`
	Browsers map[string]*browser.Browser `yaml:"browsers"`
	Rules    []*rule.Rule                `yaml:"rules"`
	Debug    bool                        `yaml:"debug"`
}

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

// New instance of Config
func New() (*Config, error) {
	c := &Config{}

	config, err := ioutil.ReadFile(fmt.Sprintf("%s/.config/choosy/config", os.Getenv("HOME")))

	if os.IsNotExist(err) {
		fmt.Println("File doesn't exist")
		errorString := "http://juancmuller.com/simplemessage/choosyerror.html?home=%s"
		exec.Command("chromium-browser", fmt.Sprintf(errorString, os.Getenv("HOME"))).Run()
		os.Exit(0)
	}

	err = yaml.Unmarshal(config, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
