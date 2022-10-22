package main

import (
	"image/color"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	windowWidth  = 600
	windowHeight = 600
	cellSize     = 10
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Cellular Zoo",
		Bounds: pixel.R(0, 0, windowWidth, windowHeight),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)

	for !win.Closed() {
		win.Clear(colornames.Snow)
		imd.Clear()

		drawRect(imd, windowWidth/2, windowHeight/2, cellSize, cellSize, colornames.Tomato)
		drawRect(imd, 20, 20, cellSize, cellSize, colornames.Wheat)

		imd.Draw(win)
		win.Update()
	}
}

func drawRect(imd *imdraw.IMDraw, x, y, w, h float64, color color.RGBA) {
	rect := pixel.R(x, y, x+w, y+h).Norm()

	imd.Color = color
	imd.Push(rect.Min, rect.Max)
	imd.Rectangle(0)
}

func main() {
	pixelgl.Run(run)
}
