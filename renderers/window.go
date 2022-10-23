package renderers

import (
	"github.com/zronev/cellular-zoo/colony"
	"github.com/zronev/cellular-zoo/config"
	"github.com/zronev/cellular-zoo/rules"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

// TODO: Split into separate struct
type WindowRenderer struct {
	win          *pixelgl.Window
	imd          *imdraw.IMDraw
	colony       *colony.Colony
	colonyDrawer *colony.ColonyDrawer
}

func (wr *WindowRenderer) Setup() error {
	cfg := pixelgl.WindowConfig{
		Title:  "Cellular Zoo",
		Bounds: pixel.R(0, 0, config.WindowWidth, config.WindowHeight),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		return err
	}

	wr.win = win
	wr.imd = imdraw.New(nil)
	wr.colony = colony.New(config.WindowHeight/config.CellSize, config.WindowWidth/config.CellSize)
	wr.colonyDrawer = colony.NewDrawer(wr.colony, config.CellSize)

	return nil
}

func (wr *WindowRenderer) Update() {
	wr.colony.NextGen(func(cell colony.Cell, neighbours int) colony.Cell {
		return colony.Cell(rules.GOL(rules.GOLCell(cell), neighbours))
	})
}

func (wr *WindowRenderer) Draw() {
	wr.colonyDrawer.Draw(wr.imd)
	wr.imd.Draw(wr.win)
	wr.win.Update()
}

func (wr *WindowRenderer) Clear() {
	wr.win.Clear(colornames.Snow)
	wr.imd.Clear()
}

func (wr *WindowRenderer) Running() bool {
	return !wr.win.Closed()
}