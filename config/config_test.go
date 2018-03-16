package config_test

import (
	"regexp"
	"testing"

	"github.com/jcmuller/picky/config"
)

var sourceString = `---
debug: true
browsers:
	foobar:
		path: foobar_path
		profile: foobar_profile %s
	barbaz:
		path: barbaz_path
		profile: barbaz_profile %s
default:
	browser: foobar
	profile: default_profile
rules:
	- uri: first_url
		browser: barbaz
		profile: first_profile
	- uri: second_uri
		browser: foobar
		profile: second_profile
`

func TestNew(t *testing.T) {
	yaml := []byte(regexp.MustCompile("\t").ReplaceAllString(sourceString, "  "))

	config, err := config.New(yaml)

	if err != nil {
		t.Errorf("Not supposed to get error: %+v", err)
	}

	if len(config.GetBrowsers()) != 2 {
		t.Error("Incorrect number of browsers")
	}

	if !config.GetDebug() {
		t.Error("Incorrect parsing of debug")
	}

	if len(config.GetRules()) != 2 {
		t.Error("Incorrect number of rules")
	}

	if config.GetDefaultRule().GetProfile() != "default_profile" {
		t.Error("Error parsing default rule's profile")
	}

	if config.GetDefaultRule().GetBrowser() != "foobar" {
		t.Error("Error parsing default rule's browser")
	}
}
