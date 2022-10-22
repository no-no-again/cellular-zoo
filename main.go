package main

import (
	"github.com/zronev/cellular-zoo/colony"
	"github.com/zronev/cellular-zoo/config"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	cellSize = 10
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Cellular Zoo",
		Bounds: pixel.R(0, 0, config.WindowWidth, config.WindowHeight),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)

	c := colony.New(config.WindowHeight/cellSize, config.WindowWidth/cellSize)
	cd := colony.NewDrawer(c, cellSize)

	for !win.Closed() {
		win.Clear(colornames.Snow)
		imd.Clear()

		cd.Draw(imd)

		imd.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
