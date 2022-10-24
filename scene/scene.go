package scene

import (
	"time"

	"github.com/zronev/cellular-zoo/config"
	"github.com/zronev/cellular-zoo/drawer"
	"github.com/zronev/cellular-zoo/renderers"
)

type Scene interface {
	Update()
	Draw(drawer drawer.Drawer)
}

type Opts struct {
	Renderer renderers.Renderer
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
	ticker := time.NewTicker(time.Second / config.FPS)

	for opts.Renderer.Running() {
		opts.Renderer.Clear()

		scene.Update()

		opts.Renderer.Draw(func(d drawer.Drawer) {
			scene.Draw(d)
		})

		<-ticker.C
	}
}
