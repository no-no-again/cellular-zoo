package scene

import (
	"log"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/zronev/cellular-zoo/config"
	"github.com/zronev/cellular-zoo/drawers"
	"golang.org/x/image/colornames"
)

type Scene interface {
	Setup()
	Update()
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
		log.Fatal("failed to create a window", err)
	}

	imd := imdraw.New(nil)
	drawer := drawers.NewWindowDrawer(imd, int(cfg.Bounds.W()), int(cfg.Bounds.H()))

	scene.Setup()

	ticker := time.NewTicker(time.Second / config.FPS)

	for !win.Closed() {
		win.Clear(colornames.Snow)
		imd.Clear()

		scene.Update()
		scene.Draw(drawer)

		imd.Draw(win)
		win.Update()

		<-ticker.C
	}
}
