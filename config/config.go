// Package config has configuration responsibilities
package config

import (
	"errors"
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
	Browsers    map[string]*browser.Browser `yaml:"browsers"`
	Debug       bool                        `yaml:"debug"`
	DefaultRule *rule.Rule                  `yaml:"default"`
	Rules       []*rule.Rule                `yaml:"rules"`
}

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

var configFileTemplate = "%s/.config/choosy/config"

// FilePath returns the config file path
func FilePath() (file string, err error) {
	file = fmt.Sprintf(configFileTemplate, os.Getenv("HOME"))
	return
}

// FileContents reads the config file
func FileContents(path string) (configFile []byte, err error) {
	configFile, err = ioutil.ReadFile(path)
	if os.IsNotExist(err) {
		errorString := "http://juancmuller.com/simplemessage/choosyerror.html?home=%s"
		err = exec.Command("chromium-browser", fmt.Sprintf(errorString, os.Getenv("HOME"))).Run()
		handle(err)

		return nil, errors.New("config not found")
	}

	handle(err)

	return
}

// New instance of Config
func New(configYaml []byte) (*Config, error) {
	c := &Config{}
	err := yaml.Unmarshal(configYaml, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}

// GetDefaultRule returns the default rule
func (c *Config) GetDefaultRule() *rule.Rule {
	return c.DefaultRule
}

// GetBrowsers returns the browsers
func (c *Config) GetBrowsers() map[string]*browser.Browser {
	return c.Browsers
}

// GetRules returns the rules
func (c *Config) GetRules() []*rule.Rule {
	return c.Rules
}

// GetDebug returns the debug flag
func (c *Config) GetDebug() bool {
	return c.Debug
}
