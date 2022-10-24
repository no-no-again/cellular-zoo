package main

import (
	"fmt"

	"github.com/zronev/cellular-zoo/config"
	"github.com/zronev/cellular-zoo/drawers"
	"github.com/zronev/cellular-zoo/renderers"
	"github.com/zronev/cellular-zoo/rule"
	"github.com/zronev/cellular-zoo/scene"
	"github.com/zronev/cellular-zoo/world"
)

type State struct {
	rule        *rule.Rule
	world       *world.World
	worldDrawer *world.Drawer
}

type Scene struct {
	state *State
}

func (s *Scene) Update() {
	s.state.world.NextGen(s.state.rule)
}

func (s *Scene) Draw(drawer drawers.Drawer) {
	s.state.worldDrawer.Draw(drawer)
}

func main() {
	const ruleString = config.DefaultRule

	sceneOpts := &scene.Opts{
		Renderer: &renderers.WindowRenderer{},
	}

	rule, err := rule.FromString(ruleString)
	if err != nil {
		panic(fmt.Sprintln("wrong rule: ", ruleString))
	}

	wld := world.New(
		config.WindowHeight/config.CellSize,
		config.WindowWidth/config.CellSize,
		rule.States(),
	)
	wldDrawer := world.NewDrawer(wld, config.CellSize)

	state := &State{rule, wld, wldDrawer}
	myScene := &Scene{state}

	scene.Run(myScene, sceneOpts)
}
