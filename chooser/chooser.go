package chooser

import (
	"fmt"
	"os/exec"

	"github.com/jcmuller/choosy/config"
)

// Chooser thingie
type Chooser struct {
	arg string
}

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

// New instance of chooser
func New(arg string) *Chooser {
	return &Chooser{arg: arg}
}

func (c *Chooser) getRule() *config.Rule {
	config, err := config.New()
	handle(err)

	for _, r := range config.Rules {
		if r.Match(c.arg) {
			return r
		}
	}

	return config.Default
}

// Call runs this thing
func (c *Chooser) Call() {
	rule := c.getRule()

	err := exec.Command("chromium-browser", fmt.Sprintf("--profile-directory=%s", rule.Profile), c.arg).Run()
	handle(err)
}
