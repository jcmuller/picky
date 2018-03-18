// Package config has configuration responsibilities
package config

import (
	"github.com/jcmuller/picky/rule"
)

// Config is a struct that holds configuration
type Config struct {
	Debug   bool
	Default *rule.Rule `yaml:"default"`
	Rules   []*rule.Rule
}

// GetDefaultRule returns the default rule
func (c *Config) GetDefaultRule() *rule.Rule {
	return c.Default
}

// GetRules returns the rules
func (c *Config) GetRules() []*rule.Rule {
	return c.Rules
}

// GetDebug returns the debug flag
func (c *Config) GetDebug() bool {
	return c.Debug
}
