package rule

import (
	"regexp"

	yaml "gopkg.in/yaml.v2"
)

// Rule defines what to do for a URL
type Rule struct {
	uri     string `yaml:"uri"`
	Profile string `yaml:"profile"`
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
	match, err := regexp.MatchString(r.uri, host)

	if err != nil {
		return false
	}

	return match
}
