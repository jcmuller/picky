package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/jcmuller/choosy/rule"

	yaml "gopkg.in/yaml.v2"
)

// Config is a struct that holds configuration
type Config struct {
	Default *rule.Rule   `yaml:"default"`
	Rules   []*rule.Rule `yaml:"rules"`
	Debug   bool         `yaml:"debug"`
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
		exec.Command("chromium-browser", fmt.Sprintf("http://juancmuller.com/simplemessage/?heading=Error&content=You need to create %s/.config/choosy/config", os.Getenv("HOME"))).Run()
		os.Exit(0)
	}

	err = yaml.Unmarshal(config, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
