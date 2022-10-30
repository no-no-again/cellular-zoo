package config

import (
	"github.com/zronev/cellular-zoo/world"
	"golang.org/x/image/colornames"
)

const (
	WindowWidth   = 600
	WindowHeight  = 700
	SpawnCapacity = 0.5
	FPS           = 10.0
	FrameRate     = 1.0 / FPS
	CellSize      = 5
	DefaultRule   = "2-3/3/2/M"
)

// TODO: support 10 colors
var DefaultPalette = world.Palette{
	colornames.Snow,
	colornames.Darkslategray,
	colornames.Tomato,
	colornames.Wheat,
	colornames.Cornflowerblue,
}

var DefaultBackground = DefaultPalette[0]
