package main

import (
	"time"

	"github.com/zronev/cellular-zoo/colony"
	"github.com/zronev/cellular-zoo/config"
	"github.com/zronev/cellular-zoo/rules"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
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

	c := colony.New(config.WindowHeight/config.CellSize, config.WindowWidth/config.CellSize)
	cd := colony.NewDrawer(c, config.CellSize)

	update := func() {
		c.NextGen(func(cell colony.Cell, neighbours int) colony.Cell {
			return colony.Cell(rules.GOL(rules.GOLCell(cell), neighbours))
		})
	}

	draw := func() {
		cd.Draw(imd)
	}

	timeStart := time.Now().UnixNano()

	for !win.Closed() {
		now := time.Now().UnixNano()
		delta := float64(now-timeStart) / 1000000000

		win.Clear(colornames.Snow)
		imd.Clear()

		if delta >= config.FrameRate {
			timeStart = now
			update()
		}

		draw()

		imd.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
