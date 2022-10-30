package main

import (
	"fmt"

	"github.com/faiface/pixel/pixelgl"
	"github.com/no-no-again/cellular-zoo/config"
	"github.com/no-no-again/cellular-zoo/drawers"
	"github.com/no-no-again/cellular-zoo/rule"
	"github.com/no-no-again/cellular-zoo/scene"
	"github.com/no-no-again/cellular-zoo/world"
)

type state struct {
	rule        *rule.Rule
	world       *world.World
	worldDrawer *world.Drawer
}

type myScene struct {
	state *state
}

func (s *myScene) Setup() {
	const ruleString = config.DefaultRule

	rule, err := rule.FromString(ruleString)
	if err != nil {
		panic(fmt.Sprintln("invalid rule: ", ruleString))
	}

	rols := config.WindowHeight / config.CellSize
	cols := config.WindowWidth / config.CellSize

	w := world.New(rols, cols, rule, config.SpawnCapacity)
	wd := world.NewDrawer(w, config.DefaultPalette, config.CellSize)

	s.state = &state{
		rule:        rule,
		world:       w,
		worldDrawer: wd,
	}
}

func (s *myScene) Update() {
	s.state.world.NextGen(s.state.rule)
}

func (s *myScene) Draw(drawer drawers.Drawer) {
	s.state.worldDrawer.Draw(drawer)
}

func (s *myScene) Input(win *pixelgl.Window) {
	if win.JustPressed(pixelgl.KeyS) {
		s.state.world.Spawn(s.state.rule, config.SpawnCapacity)
	}

	if win.JustPressed(pixelgl.KeyC) {
		s.state.world.Clear()
	}
}

func main() {
	scene.Run(&myScene{})
}
