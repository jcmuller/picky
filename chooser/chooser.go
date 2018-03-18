// Package chooser chooses what browser to use depending on a rule and URI
package chooser

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/jcmuller/picky/rule"
)

type cfg interface {
	GetRules() []*rule.Rule
	GetDefaultRule() *rule.Rule
	GetDebug() bool
}

// Chooser thingie
type Chooser struct {
	arg    string
	config cfg
}

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

// New instance of chooser
func New(config cfg, arg string) *Chooser {
	return &Chooser{
		arg:    arg,
		config: config,
	}
}

// GetRule returns the rule that matches the argument passed in
func (c *Chooser) GetRule() *rule.Rule {
	config := c.config

	for _, r := range config.GetRules() {
		if r.Match(c.arg) {
			return r
		}
	}

	return config.GetDefaultRule()
}

// Call runs this thing
func (c *Chooser) Call() {
	rule := c.GetRule()
	command := rule.GetCommand(c.arg)

	if c.config.GetDebug() {
		fmt.Fprintf(os.Stderr, "%s\n", command)
		return
	}

	err := exec.Command(command[0], command[1:]...).Run()
	handle(err)
}
