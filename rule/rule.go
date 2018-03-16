// Package rule wraps heuristics to choose browsers
package rule

import (
	"fmt"
	"regexp"
)

// Rule defines what to do for a URL
type Rule struct {
	Base    string   `yaml:"base"`
	Profile string   `yaml:"profile"`
	Args    string   `yaml:"args"`
	URIs    []string `yaml:"uris"`
}

// GetCommand returns the browser key for the rule
func (r *Rule) GetCommand(uri string) [3]string {
	return [3]string{
		r.Base,
		fmt.Sprintf(r.Profile, r.Args),
		uri,
	}
}

// Match matches
func (r *Rule) Match(host string) bool {
	for _, uri := range r.URIs {
		re := regexp.MustCompile(uri)

		if re.MatchString(host) {
			return true
		}
	}

	return false
}
