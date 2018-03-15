package rule

import (
	"regexp"

	yaml "gopkg.in/yaml.v2"
)

// Rule defines what to do for a URL
type Rule struct {
	URI     string `yaml:"uri"`
	Profile string `yaml:"profile"`
	Browser string `yaml:"browser"`
}

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

func (r *Rule) String() string {
	b, err := yaml.Marshal(r)
	handle(err)
	return string(b)
}

// Match matches
func (r *Rule) Match(host string) bool {
	match, err := regexp.MatchString(r.URI, host)

	if err != nil {
		return false
	}

	return match
}
