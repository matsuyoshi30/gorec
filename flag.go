package main

import (
	"flag"
	"fmt"
)

const USAGE = `
NAME:
   gorec - Simple Recorder in Go
USAGE:
   gorec
VERSION:
   0.1.0
GLOBAL OPTIONS:
  -o              set output filename
  -l              use login shell
  -b              use bold font style
  --help, -h      show help
`

type Option struct {
	output string
	login  bool
	bold   bool
}

func parseFlag(args ...string) *Option {
	fs := flag.NewFlagSet("gorec", flag.ExitOnError)
	fs.Usage = func() {
		fmt.Println(USAGE)
	}

	out := fs.String("o", "output", "set output filename")
	l := fs.Bool("l", false, "use login shell")
	b := fs.Bool("b", false, "use bold font style")
	fs.Parse(args)

	return &Option{output: *out, login: *l, bold: *b}
}
