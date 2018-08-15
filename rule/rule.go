// Package rule wraps heuristics to choose browsers
package rule

import (
	"regexp"
)

// Rule defines what to do for a URL
type Rule struct {
	Label   string
	Command string
	Args    []string
	URIs    []string
}

// GetCommand returns the browser key for the rule
func (r *Rule) GetCommand() (command string, args []string) {
	command = r.Command
	args = r.Args

	return
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
