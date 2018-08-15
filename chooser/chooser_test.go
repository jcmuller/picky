package chooser_test

import (
	"testing"

	"github.com/jcmuller/picky/chooser"
	"github.com/jcmuller/picky/rule"
)

type mockConfig struct{}

var defaultRule = &rule.Rule{
	Command: "command",
	Args:    []string{"--profile=Some Profile"},
}

var otherRule = &rule.Rule{
	Command: "otherCommand",
	Args:    []string{"-P=Another Profile"},
	URIs:    []string{"hotmail.com", "yahoo.com"},
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
	c := chooser.New(&mockConfig{}, "hotmail.com")

	rule := c.GetRule()

	if rule != otherRule {
		t.Errorf("Rules don't match:\n  expected: %+v\n    actual: %+v", otherRule, rule)

	}
}

func TestCall(t *testing.T) {
	c := chooser.New(&mockConfig{}, "some awesome argument")

	c.Call()

	if c == nil {
		t.Errorf("Something went wrong")
	}
}
