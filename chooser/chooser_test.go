package chooser_test

import (
	"testing"

	"github.com/jcmuller/choosy/browser"
	"github.com/jcmuller/choosy/chooser"
	"github.com/jcmuller/choosy/rule"
)

type mockConfig struct{}

var defaultRule = &rule.Rule{
	Profile: "default_profile",
	Browser: "default_browser",
}

var otherRule = &rule.Rule{
	URI:     "other_uri",
	Profile: "profile",
	Browser: "default_browser",
}

func (c *mockConfig) GetBrowsers() map[string]*browser.Browser {
	a := make(map[string]*browser.Browser)

	b := &browser.Browser{
		Path:    "path",
		Profile: "foobar %s",
	}

	a["default_browser"] = b

	return a
}

func (c *mockConfig) GetRules() []*rule.Rule {
	return []*rule.Rule{otherRule}
}

func (c *mockConfig) GetDefaultRule() *rule.Rule {
	return defaultRule
}

func (c *mockConfig) GetDebug() bool {
	return true
}

func TestGetDefaultRule(t *testing.T) {
	c := chooser.New(&mockConfig{}, "some awesome argument")

	rule := c.GetRule()

	if rule != defaultRule {
		t.Errorf("Rules don't match:\n  expected: %+v\n    actual: %+v", defaultRule, rule)
	}
}

func TestGetRule(t *testing.T) {
	c := chooser.New(&mockConfig{}, "other_uri")

	rule := c.GetRule()

	if rule != otherRule {
		t.Errorf("Rules don't match:\n  expected: %+v\n    actual: %+v", otherRule, rule)
	}
}

func TestCall(t *testing.T) {
	c := chooser.New(&mockConfig{}, "some awesome argument")

	c.Call()

	if c == nil {
		t.Errorf("You fucked up")
	}
}
