package renderers

import "github.com/zronev/cellular-zoo/drawer"

type Renderer interface {
	Setup() error
	Run(run func())
	Draw(draw func(d drawer.Drawer))
	Clear()
	Running() bool
}
