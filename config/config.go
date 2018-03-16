// Package config has configuration responsibilities
package config

import (
	"github.com/jcmuller/picky/browser"
	"github.com/jcmuller/picky/rule"

	yaml "gopkg.in/yaml.v2"
)

// Config is a struct that holds configuration
type Config struct {
	Browsers    map[string]*browser.Browser `yaml:"browsers"`
	Debug       bool                        `yaml:"debug"`
	DefaultRule *rule.Rule                  `yaml:"default"`
	Rules       []*rule.Rule                `yaml:"rules"`
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
