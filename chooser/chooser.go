// Package chooser chooses what browser to use depending on a rule and URI
package chooser

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/jcmuller/picky/browser"
	"github.com/jcmuller/picky/rule"
)

type cfg interface {
	GetBrowsers() map[string]*browser.Browser
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

func log(command string) {
	f, err := os.OpenFile("/tmp/input", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	handle(err)
	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf("%q\n", command))
	handle(err)
}

// Call runs this thing
func (c *Chooser) Call() {
	rule := c.GetRule()

	b := c.config.GetBrowsers()[rule.GetBrowser()]
	command := b.GetCommand(rule, c.arg)

	if c.config.GetDebug() {
		log(fmt.Sprintf("%v", command))
		return
	}

	err := exec.Command(command[0], command[1:]...).Run()
	handle(err)
}
