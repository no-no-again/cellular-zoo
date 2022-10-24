package renderers

import (
	"github.com/zronev/cellular-zoo/config"
	"github.com/zronev/cellular-zoo/drawers"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

type WindowRenderer struct {
	drawer drawers.Drawer
	win    *pixelgl.Window
	imd    *imdraw.IMDraw
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
	wr.drawer = drawers.NewWindowDrawer(wr.imd)

	return nil
}

func (wr *WindowRenderer) Run(run func()) {
	pixelgl.Run(run)
}

func (wr *WindowRenderer) Draw(draw func(d drawers.Drawer)) {
	draw(wr.drawer)
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
