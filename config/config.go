package config

import (
	"io/ioutil"

	"github.com/jcmuller/choosy/rule"

	yaml "gopkg.in/yaml.v2"
)

// Config is a struct that holds configuration
type Config struct {
	Default *rule.Rule   `yaml:"default"`
	Rules   []*rule.Rule `yaml:"rules"`
}

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

// New instance of Config
func New() (*Config, error) {
	c := &Config{}

	config, err := ioutil.ReadFile("/home/jcmuller/.config/choosy/config")
	handle(err)

	err = yaml.Unmarshal(config, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
