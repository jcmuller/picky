// Package browser knows how to deal with browsers
package browser

import (
	"fmt"
)

// Browser represents a browser definition
type Browser struct {
	Path    string `yaml:"path"`
	Profile string `yaml:"profile"`
}

type rule interface {
	Match(string) bool
	GetProfile() string
}

// GetCommand returns the command to launch the browser for a given rule and an URI
func (b *Browser) GetCommand(r rule, url string) [3]string {
	return [3]string{b.Path, fmt.Sprintf(b.Profile, r.GetProfile()), url}
}
