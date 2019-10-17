package main

import (
	"fmt"
	"image"
	"os"
	"time"
)

type TW struct {
	output         string
	outputFile     *os.File // debug
	timestampStart time.Time
	recdata        []*RecData
	images         []*image.Paletted
	delays         []int
	bold           bool
}

func main() {
	o := parseFlag(os.Args[1:]...)

	tw, err := Rec(o.output, o.login, o.bold)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
	}
	tw.Generate()

	fmt.Println(tw.output+".gif", "CREATED!!")
}
