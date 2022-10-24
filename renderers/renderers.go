package renderers

import "github.com/zronev/cellular-zoo/drawers"

type Renderer interface {
	Setup() error
	Run(run func())
	Draw(draw func(d drawers.Drawer))
	Clear()
	Running() bool
}
