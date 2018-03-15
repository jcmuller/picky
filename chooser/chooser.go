package chooser

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/jcmuller/choosy/config"
	"github.com/jcmuller/choosy/rule"
)

// Chooser thingie
type Chooser struct {
	arg    string
	config *config.Config
}

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

// New instance of chooser
func New(arg string) *Chooser {
	config, err := config.New()

	handle(err)

	return &Chooser{
		arg:    arg,
		config: config,
	}
}

func (c *Chooser) getRule() *rule.Rule {
	config := c.config

	for _, r := range config.Rules {
		if r.Match(c.arg) {
			return r
		}
	}

	return config.Default
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
	rule := c.getRule()

	b := c.config.Browsers[rule.Browser]
	command := b.GetCommand(rule, c.arg)

	if c.config.Debug {
		log(strings.Join(command, " "))
		return
	}

	err := exec.Command(command[0], command[1:]...).Run()
	handle(err)
}
