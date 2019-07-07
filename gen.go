package main

import (
	"image/gif"
	"os"
)

func (t *TW) Generate() error {
	err := t.Play()
	if err != nil {
		return err
	}

	file, err := os.Create(t.output + ".gif")
	if err != nil {
		return err
	}
	defer os.Remove(t.output)
	defer file.Close()

	g := gif.GIF{
		Image: t.images,
		Delay: t.delays,
	}

	err = gif.EncodeAll(file, &g)
	if err != nil {
		return err
	}

	return nil
}
