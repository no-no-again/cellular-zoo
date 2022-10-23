package scene

import (
	"time"

	"github.com/zronev/cellular-zoo/drawer"
	"github.com/zronev/cellular-zoo/renderers"
)

type Scene interface {
	Update()
	Draw(drawer drawer.Drawer)
}

type Opts struct {
	Renderer  renderers.Renderer
	FrameRate float64
}

func Run(scene Scene, opts *Opts) {
	opts.Renderer.Run(func() {
		err := opts.Renderer.Setup()
		if err != nil {
			panic(err)
		}
		loop(scene, opts)
	})
}

func loop(scene Scene, opts *Opts) {
	timeStart := time.Now().UnixNano()

	for opts.Renderer.Running() {
		now := time.Now().UnixNano()
		delta := float64(now-timeStart) / 1_000_000_000

		opts.Renderer.Clear()
		if delta >= opts.FrameRate {
			timeStart = now
			scene.Update()
		}
		opts.Renderer.Draw(func(d drawer.Drawer) {
			scene.Draw(d)
		})
	}
}
