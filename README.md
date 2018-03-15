# Choosy

This program lets you open URIs using different browsers (or profiles) following a set of rules.

It is driven by a configuration file located at `$HOME/.config/choosy/config`.

## Example config
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
