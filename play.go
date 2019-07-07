package main

import (
	"encoding/base64"
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"time"

	"github.com/mattn/go-libvterm"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

func (t *TW) Play() error {
	vt := vterm.New(25, 80)
	defer vt.Close()

	screen := vt.ObtainScreen()
	screen.Reset(true)

	var prevTs int
	for _, rd := range t.recdata {
		var diff int

		if prevTs == 0 {
			clearScreen()
		} else {
			diff = rd.ts - prevTs
		}
		prevTs = rd.ts
		time.Sleep(time.Millisecond * time.Duration(diff) / 5000)

		str, err := base64.StdEncoding.DecodeString(rd.encdata)
		if err != nil {
			return err
		}

		_, err = vt.Write(str)
		if err != nil {
			return err
		}
		screen.Flush()

		rows, cols := vt.Size()
		img := image.NewRGBA(image.Rect(0, 0, cols*7, rows*13))
		draw.Draw(img, img.Bounds(), &image.Uniform{color.Black}, image.ZP, draw.Src)

		for row := 0; row < rows; row++ {
			for col := 0; col < cols; col++ {
				cell, err := screen.GetCellAt(row, col)
				if err != nil {
					return err
				}
				chars := cell.Chars()
				if len(chars) > 0 && chars[0] != 0 {
					drawChar(img, (col+1)*7, (row+1)*13, cell.Fg(), string(chars))
				}
			}
		}

		palettedImage := image.NewPaletted(img.Bounds(), palette.Plan9)
		draw.FloydSteinberg.Draw(palettedImage, img.Bounds(), img, image.ZP)

		t.images = append(t.images, palettedImage)
		t.delays = append(t.delays, diff)
	}

	return nil
}

func drawChar(img *image.RGBA, x, y int, c color.Color, text string) {
	point := fixed.Point26_6{fixed.Int26_6(x * 64), fixed.Int26_6(y * 64)}
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(c),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(text)
}
