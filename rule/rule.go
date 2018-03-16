// Package rule wraps heuristics to choose browsers
package rule

import (
	"regexp"
)

// Rule defines what to do for a URL
type Rule struct {
	Browser string `yaml:"browser"`
	Profile string `yaml:"profile"`
	URI     string `yaml:"uri"`
}

// GetProfile returns the profile
func (r *Rule) GetProfile() string {
	return r.Profile
}

// GetBrowser returns the browser key for the rule
func (r *Rule) GetBrowser() string {
	return r.Browser
}

// Match matches
func (r *Rule) Match(host string) bool {
	match, err := regexp.MatchString(r.URI, host)

	if err != nil {
		return false
	}

	return match
}
