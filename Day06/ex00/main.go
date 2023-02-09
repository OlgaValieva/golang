package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

func main() {
	width := 300
	height := 300

	aqua := color.RGBA{R: 0, G: 255, B: 255, A: 0xff}
	yellow := color.RGBA{R: 255, G: 255, B: 0, A: 0xff}
	blue := color.RGBA{R: 0, G: 0, B: 255, A: 0xff}

	upLeft := image.Point{}
	lowRight := image.Point{X: width, Y: height}

	img := image.NewRGBA(image.Rectangle{Min: upLeft, Max: lowRight})

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			switch {
			case x < width/3 && y >= 2*height/3:
				img.Set(x, y, blue)
			case x >= 2*width/3 && y < height/3:
				img.Set(x, y, blue)
			case x < width/3 && y < height/3:
				img.Set(x, y, blue)
			case x >= 2*width/3 && y >= 2*height/3:
				img.Set(x, y, blue)
			case x >= width/3 && x < 2*width/3 && y >= height/3 && y < 2*height/3:
				img.Set(x, y, yellow)
			default:
				img.Set(x, y, aqua)
			}
		}
	}

	addLabel(img, 2*width/5, height/2, " carys:)")

	f, er := os.Create("amazing_logo.png")
	if er != nil {
		fmt.Println("Error - no create", er)
		os.Exit(1)
	}
	defer f.Close()

	err := png.Encode(f, img)

	if err != nil {
		return
	}
}

func addLabel(img *image.RGBA, x, y int, label string) {
	col := color.RGBA{R: 255, G: 0, B: 255, A: 0xff}
	point := fixed.Point26_6{X: fixed.I(x), Y: fixed.I(y)}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(label)
}
