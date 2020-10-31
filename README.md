# ctap
[![Go](https://github.com/Sean0628/ctap/workflows/Go/badge.svg?branch=main)](https://github.com/Sean0628/ctap/actions?query=workflow%3AGo)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://github.com/Sean0628/ctap/blob/main/LICENSE.md)

## Overview
`ctap` is a CLI **c**ron**ta**b **p**arser written in Go.

This allows you to convert cron expression into human readable description, and output into the specified format(csv/markdown).
The command is based on [hcron](https://github.com/lnquy/cron#hcron) created by [Quy Le](https://github.com/lnquy).

## Usage
```sh
$ ctap -h
ctap is a CLI crontab parser written in Go.

Usage:
  ctap [flags]

Flags:
      --24-hour             Prints out description in 24 hour time format.
  -d, --dow-starts-at-one   Is day of the week starts at 1 (Monday-Sunday: 1-7).
  -f, --format string       Prints out in the specified format. options: csv, markdown
  -h, --help                help for ctap
  -i, --input string        Path to a crontab file to be read from.
  -l, --locale string       Prints out description in the specified locale. (default "en")
  -o, --output string       Path to an output file. If this option is not set, the results are printed out to standard output.
  -s, --sort                Prints out cron expressions in ascending order.
  -v, --verbose             Prints out description in verbose format.
  -V, --version             Prints out version information.
```

## Installation
### with go get
```sh
$ go get -u -v github.com/Sean0628/ctap
```

## Copyright
Copyright (c) 2020 Sho ITO. See [LICENSE.md](https://github.com/Sean0628/ctap/blob/main/LICENSE.md) for further details.
