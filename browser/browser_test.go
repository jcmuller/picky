package browser_test

import (
	"testing"

	"github.com/jcmuller/picky/browser"
)

type mockRule struct {
}

func (r *mockRule) GetProfile() string {
	return "profile_name"
}

func (r *mockRule) Match(s string) bool {
	return false
}

func TestGetCommand(t *testing.T) {
	browser := &browser.Browser{
		Path:    "browser_path",
		Profile: "profile_option %s",
	}

	rule := &mockRule{}

	actual := browser.GetCommand(rule, "uri")
	expected := [3]string{"browser_path", "profile_option profile_name", "uri"}

	if actual != expected {
		t.Errorf("Command is incorrect:\n  expected: %+v\n  actual:   %+v", expected, actual)
	}
}

func BenchmarkGetCommand(b *testing.B) {
	browser := &browser.Browser{
		Path:    "browser_path",
		Profile: "profile_option %s",
	}

	rule := &mockRule{}

	for n := 0; n < b.N; n++ {
		browser.GetCommand(rule, "uri")
	}
}
