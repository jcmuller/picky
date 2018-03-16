package rule_test

import (
	"testing"

	"github.com/jcmuller/picky/rule"
)

var matchTests = []struct {
	uri      string
	expected bool
}{
	{"https://uri/foo/bar", true},
	{"https://uri/zoo/bar", true},
	{"ttps://uri/foo/bar", false},
	{"http://uri/foo/bar", true},
	{"https://uri/foo/barista", true},
	{"https://uri/foo/bar/bafo", true},
}

func TestMatch(t *testing.T) {
	r := &rule.Rule{
		Base:    "browser",
		Profile: "--profile %s",
		Args:    "foobar",
		URIs:    []string{"^https?://uri/[fz]oo/bar"},
	}

	for _, tt := range matchTests {
		if r.Match(tt.uri) != tt.expected {
			t.Errorf("URI %s should have matched %s", tt.uri, r.URIs)
		}
	}
}

func TestCommand(t *testing.T) {
	r := &rule.Rule{
		Base:    "browser",
		Profile: "--profile %s",
		Args:    "some profile",
	}

	actual := r.GetCommand("foobar")

	expected := [3]string{"browser", "--profile some profile", "foobar"}

	if actual != expected {
		t.Errorf("Error:\n  expected: %+v\n    actual: %+v", expected, actual)
	}
}
