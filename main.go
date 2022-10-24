package main

import (
	"fmt"

	"github.com/zronev/cellular-zoo/colony"
	"github.com/zronev/cellular-zoo/config"
	"github.com/zronev/cellular-zoo/drawer"
	"github.com/zronev/cellular-zoo/renderers"
	"github.com/zronev/cellular-zoo/rule"
	"github.com/zronev/cellular-zoo/scene"
)

type State struct {
	rule         *rule.Rule
	colony       *colony.Colony
	colonyDrawer *colony.Drawer
}

type Scene struct {
	state *State
}

func (s *Scene) Update() {
	s.state.colony.NextGen(s.state.rule)
}

func (s *Scene) Draw(drawer drawer.Drawer) {
	s.state.colonyDrawer.Draw(drawer)
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

	col := colony.New(
		config.WindowHeight/config.CellSize,
		config.WindowWidth/config.CellSize,
		rule.States(),
	)
	colDrawer := colony.NewDrawer(col, config.CellSize)

	state := &State{rule, col, colDrawer}
	myScene := &Scene{state}

	scene.Run(myScene, sceneOpts)
}
