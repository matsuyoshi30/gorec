package main

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"

	"github.com/creack/pty"
	"golang.org/x/crypto/ssh/terminal"
)

type RecData struct {
	ts      int
	data    []byte
	encdata string
}

func Rec(dstFile string, login bool) (*TW, error) {
	var c *exec.Cmd
	if login {
		c = exec.Command(os.Getenv("SHELL"), "-i", "-l")
	} else {
		c = exec.Command(os.Getenv("SHELL"), "-i")
	}

	tm, err := pty.Start(c)
	if err != nil {
		return nil, err
	}
	defer tm.Close()

	clearScreen()

	oldState, err := terminal.MakeRaw(0)
	if err != nil {
		return nil, err
	}
	defer terminal.Restore(0, oldState)

	tw := &TW{
		output: dstFile,
	}

	tw.outputFile, err = os.Create(tw.output)
	if err != nil {
		return nil, err
	}
	tw.timestampStart = time.Now()

	allWriter := io.MultiWriter(os.Stdout, tw)

	go func() {
		io.Copy(allWriter, tm)
	}()
	go func() {
		io.Copy(tm, os.Stdin)
	}()
	c.Wait()

	return tw, nil
}

func (t *TW) Write(data []byte) (n int, err error) {
	t.writeData(data)
	return len(data), err
}

func (t *TW) writeData(data []byte) {
	ts := int(time.Since(t.timestampStart).Seconds() * 100)

	r := &RecData{
		ts:   ts,
		data: data,
	}
	t.recdata = append(t.recdata, r)
	r.encdata = base64.StdEncoding.EncodeToString(r.data)

	fmt.Fprintf(t.outputFile, `{
			timeval: %d,
			data: "%s",
			encdata: "%s"
		}`, r.ts, r.data, r.encdata)
}
