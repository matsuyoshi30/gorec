package main

import (
	"fmt"
	"os"
	"os/exec"
)

func clearScreen() error {
	b, err := exec.Command("clear").Output()
	if err != nil {
		return err
	}
	fmt.Fprintf(os.Stdout, "%s", b)
	return nil
}
