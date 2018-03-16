# Picky

[![Build Status](https://travis-ci.org/jcmuller/picky.svg?branch=master)](https://travis-ci.org/jcmuller/picky)
[![GoDoc](https://godoc.org/github.com/jcmuller/picky?status.svg)](https://godoc.org/github.com/jcmuller/picky)
[![Go Report Card](https://goreportcard.com/badge/github.com/jcmuller/picky)](https://goreportcard.com/report/github.com/jcmuller/picky)
[![codecov](https://codecov.io/gh/jcmuller/picky/branch/master/graph/badge.svg)](https://codecov.io/gh/jcmuller/picky)

This program lets you open URIs using different browsers (or profiles) following a set of rules.

**Table of Contents**
- [Example config](#example-config)
- [Installation](#installation)
- [Usage](#usage)

## Example config
Picky is driven by a configuration file located at `$HOME/.config/picky/config`.

```yaml
---
default:
  profile: Default
  browser: chromium
browsers:
  chromium:
    path: chromium-browser
    profile: --profile-directory=%s
  firefox:
    path: firefox
    profile: -P %
rules:
  - uri: https?://google.com/foo
    profile: Profile 1
    browser: firefox
  - uri: http://yahoo.com/bar
    profile: Profile 2
    browser: chromium
```

## Installation
```bash
$ go get -u github.com/jcmuller/picky
```

## Usage
```bash
$ picky SOME_URI
```