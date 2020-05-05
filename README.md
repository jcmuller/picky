# Picky

[![Build Status](https://travis-ci.org/jcmuller/picky.svg?branch=master)](https://travis-ci.org/jcmuller/picky)
[![Go Report Card](https://goreportcard.com/badge/github.com/jcmuller/picky)](https://goreportcard.com/report/github.com/jcmuller/picky)
[![GoDoc](https://godoc.org/github.com/jcmuller/picky?status.svg)](https://godoc.org/github.com/jcmuller/picky)
[![Maintainability](https://api.codeclimate.com/v1/badges/b3da22424ebf6d92f378/maintainability)](https://codeclimate.com/github/jcmuller/picky/maintainability)
[![codecov](https://codecov.io/gh/jcmuller/picky/branch/master/graph/badge.svg)](https://codecov.io/gh/jcmuller/picky)

This program lets you open URIs using different browsers (or profiles) following a set of rules.

**Table of Contents**
- [Example config](#example-config)
- [Installation](#installation)
- [Usage](#usage)

## Example config
Picky is driven by a configuration file located at `$HOME/.config/picky/config`.

```yaml
debug: true
default: &default
  command: chromium-browser
  args: ["--profile-directory=Default Profile"]
rules:
  - <<: *default
    args: First Profile
    uris:
    - hotmail.com
    - gmail.com
  - <<: *default
    args: Second Profile
    uris:
    - (cnn|nyt).com
```

## Installation
```bash
$ go get -u github.com/jcmuller/picky

$ mkdir -p ~/.local/share/applications; cp picky.desktop ~/.local/share/applications
$ sudo update-alternatives --install /usr/bin/picky x-www-browser /usr/local/bin/picky 100
$ sudo update-alternatives --install /usr/bin/picky-gnome gnome-www-browser  /usr/local/bin/picky 100
$ xdg-settings set default-web-browser picky.desktop
```

## Usage
```bash
$ picky SOME_URI
```


