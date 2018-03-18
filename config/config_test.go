package config_test

import (
	"testing"

	"github.com/jcmuller/picky/config"
	"github.com/jcmuller/picky/rule"
)

func TestGetDebug(t *testing.T) {
	c := &config.Config{Debug: true}

	if !c.GetDebug() {
		t.Error("Did not get debug")
	}
}

func TestGetDefault(t *testing.T) {
	c := &config.Config{Default: &rule.Rule{}}
	r := c.GetDefaultRule()

	if r == nil {
		t.Error("Did not get default rule")
	}
}

func TestGetRules(t *testing.T) {
	c := &config.Config{Rules: []*rule.Rule{&rule.Rule{}, &rule.Rule{}}}
	r := c.GetRules()

	if len(r) != 2 {
		t.Error("Error getting rules")
	}
}
