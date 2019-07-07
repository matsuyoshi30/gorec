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
  --help, -h      show help
`

type Option struct {
	output string
	login  bool
}

func parseFlag(args ...string) *Option {
	fs := flag.NewFlagSet("gorec", flag.ExitOnError)
	fs.Usage = func() {
		fmt.Println(USAGE)
	}

	out := fs.String("o", "output", "set output filename")
	l := fs.Bool("l", false, "use login shell")
	fs.Parse(args)

	return &Option{output: *out, login: *l}
}
