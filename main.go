// Package main is the entrypoint
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jcmuller/picky/chooser"
	"github.com/jcmuller/picky/config"
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	usageString = `
Configuration file missing.

You need to set create a configuration file in /etc/picky or %s/.config/picky
named config.yaml with contents similar to:

---
default:
	command: chromium-browser
	args: ["--profile-directory, "Default Profile"]
rules:
	- command: chromium-browser
		args: ["--profile-directory", "First Profile"]
		uris:
			- hotmail.com
			- gmail.com
	- command: chromium-browser
		args: ["--profile-directory", "Second Profile"]
		uris:
			- (cnn|nyt).com

Picky also supports JSON and toml configuration files.
`
)

func onFileError() {
	errorString := "http://juancmuller.com/simplemessage/pickyerror.html?home=%s"
	err := open.Run(fmt.Sprintf(errorString, os.Getenv("HOME")))
	fmt.Fprintf(os.Stderr, usageString, os.Getenv("HOME"))
	if err != nil {
		panic(err)
	}
}

func printHelp() {
	fmt.Printf("Usage: %s [options] URI\n", os.Args[0])
	pflag.PrintDefaults()
}

func setupConfigFileAndFlags() (args []string, err error) {
	pflag.BoolP("debug", "d", false, "Set to print out what action we would take")
	pflag.BoolP("help", "h", false, "Show help menu")
	config := pflag.StringP("config", "c", "", "Set configuration file")
	pflag.Parse()
	args = pflag.Args()

	if *config != "" {
		viper.SetConfigFile(*config)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath("$HOME/.config/picky")
		viper.AddConfigPath("/etc/picky")
	}

	viper.SetEnvPrefix("PICKY")
	viper.AutomaticEnv()

	err = viper.BindPFlags(pflag.CommandLine)
	handle(err)

	err = viper.ReadInConfig()

	return
}

func main() {
	args, err := setupConfigFileAndFlags()

	if len(args) == 0 || viper.GetBool("help") {
		printHelp()
		os.Exit(1)
	}

	if err != nil {
		onFileError()
		os.Exit(1)
	}

	var config *config.Config
	err = viper.Unmarshal(&config)
	handle(err, "Fatal error parsing config: %s\n")

	arg := strings.Join(args, " ")

	chooser.New(config, arg).Call()
}

func handle(err error, args ...string) {
	if err != nil {
		if len(args) > 0 {
			panic(fmt.Errorf(args[0], err))
		}

		panic(err)
	}
}
