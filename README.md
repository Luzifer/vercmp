[![Go Report Card](https://goreportcard.com/badge/github.com/Luzifer/vercmp)](https://goreportcard.com/report/github.com/Luzifer/vercmp)
![](https://badges.fyi/github/license/Luzifer/vercmp)
![](https://badges.fyi/github/downloads/Luzifer/vercmp)
![](https://badges.fyi/github/latest-release/Luzifer/vercmp)
![](https://knut.in/project-status/vercmp)

# Luzifer / vercmp

`vercmp` resembles [`vercmp`](https://www.archlinux.org/pacman/vercmp.8.html) utility from [Archlinux](https://www.archlinux.org/) with two small differences: It runs on systems without `libalpm.so` and it strictly uses [SemVer](https://semver.org/) - otherwise it should be fully interchangeable.

## Usage

```console
# vercmp

Compare version numbers using SemVer version logic

Usage: vercmp <ver1> <ver2>

Output values:
  < 0 : if ver1 < ver2
    0 : if ver1 == ver2
  > 0 : if ver1 > ver2

# vercmp 7.6.0 7.4.1
1
```
