package rule_test

import (
	"reflect"
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
		Command: "browser",
		Args:    []string{"--profile", "foobar"},
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
		Command: "browser",
		Args:    []string{"--profile", "some profile"},
	}

	command, args := r.GetCommand()

	expected := []string{"--profile", "some profile"}

	if command != "browser" {
		t.Errorf("Error:\n  expected: %+v\n    actual: %+v", "browser", command)
	}

	if !reflect.DeepEqual(args, expected) {
		t.Errorf("Error:\n  expected: %+v\n    actual: %+v", expected, args)
	}
}
