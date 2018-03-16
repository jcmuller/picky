package config_test

import (
	"regexp"
	"testing"

	"github.com/jcmuller/picky/config"
)

var sourceString = `---
debug: true
default: &default
	base: chromium-browser
	profile: --profile-directory=%s
	args: Default Profile
rules:
	- <<: *default
		args: First Profile
		uris:
			- hotmail.com
			- gmail.com
	- <<: *default
		args: Second Profile
		uris:
			- (cnn|nyt).com
`

func TestNew(t *testing.T) {
	yaml := []byte(regexp.MustCompile("\t").ReplaceAllString(sourceString, "  "))

	config, err := config.New(yaml)

	if err != nil {
		t.Errorf("Not supposed to get error: %+v", err)
	}

	if !config.GetDebug() {
		t.Error("Incorrect parsing of debug")
	}

	if len(config.GetRules()) != 2 {
		t.Error("Incorrect number of rules")
	}

	if config.GetDefaultRule().Base != "chromium-browser" {
		t.Error("Error parsing default rule's profile")
	}

	if config.GetDefaultRule().Args != "Default Profile" {
		t.Error("Error parsing default rule's browser")
	}
}
