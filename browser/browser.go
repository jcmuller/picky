package browser

import (
	"fmt"

	"github.com/jcmuller/choosy/rule"
)

type Browser struct {
	Key     string `yaml:"key"`
	Path    string `yaml:"path"`
	Profile string `yaml:"profile"`
}

func (b *Browser) GetCommand(r *rule.Rule, url string) []string {
	return []string{b.Path, fmt.Sprintf(b.Profile, r.Profile), url}
}
