package loop

import (
	"time"

	"github.com/zronev/cellular-zoo/config"
	"github.com/zronev/cellular-zoo/renderers"
)

func Loop(r renderers.Renderer) {
	timeStart := time.Now().UnixNano()

	for r.Running() {
		now := time.Now().UnixNano()
		delta := float64(now-timeStart) / 1_000_000_000

		r.Clear()
		if delta >= config.FrameRate {
			timeStart = now
			r.Update()
		}
		r.Draw()
	}
}
