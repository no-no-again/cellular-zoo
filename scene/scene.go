package scene

import (
	"fmt"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/zronev/cellular-zoo/config"
	"github.com/zronev/cellular-zoo/drawers"
)

type Scene interface {
	Setup()
	Update()
	Input(win *pixelgl.Window)
	Draw(drawer drawers.Drawer)
}

func Run(scene Scene) {
	pixelgl.Run(func() { run(scene) })
}

func run(scene Scene) {
	cfg := pixelgl.WindowConfig{
		Title:  "Cellular Zoo",
		Bounds: pixel.R(0, 0, config.WindowWidth, config.WindowHeight),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(fmt.Sprintf("failed to create a window: %v", err))
	}

	imd := imdraw.New(nil)
	drawer := drawers.NewWindowDrawer(imd, int(cfg.Bounds.W()), int(cfg.Bounds.H()))

	scene.Setup()

	last := time.Now().UnixNano()
	for !win.Closed() {
		now := time.Now().UnixNano()
		dt := float64(now-last) / 1_000_000_000

		scene.Input(win)

		win.Clear(config.DefaultBackground)
		imd.Clear()

		if dt >= config.FrameRate {
			last = now
			scene.Update()
		}

		scene.Draw(drawer)
		imd.Draw(win)
		win.Update()
	}
}
