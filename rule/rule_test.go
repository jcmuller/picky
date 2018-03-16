package rule_test

import (
	"testing"

	"github.com/jcmuller/choosy/rule"
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
		Browser: "browser",
		Profile: "profile",
		URI:     "^https?://uri/[fz]oo/bar",
	}

	for _, tt := range matchTests {
		if r.Match(tt.uri) != tt.expected {
			t.Errorf("URI %s should have matched %s", tt.uri, r.URI)
		}
	}

	if r.GetBrowser() != "browser" {
		t.Error("incorrect browser")
	}

	if r.GetProfile() != "profile" {
		t.Error("incorrect profile")
	}
}
