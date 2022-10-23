package renderers

type Renderer interface {
	Setup() error
	Update()
	Draw()
	Clear()
	Running() bool
}
