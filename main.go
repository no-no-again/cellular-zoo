package main

import (
	"fmt"

	"github.com/zronev/cellular-zoo/config"
	"github.com/zronev/cellular-zoo/drawers"
	"github.com/zronev/cellular-zoo/rule"
	"github.com/zronev/cellular-zoo/scene"
	"github.com/zronev/cellular-zoo/world"
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

	w := world.New(
		rols,
		cols,
		rule.States(),
		config.SpawnCapacity,
	)
	wd := world.NewDrawer(w, config.DefaultPalette, config.CellSize)

	s.state = &state{rule, w, wd}
}

func (s *myScene) Update() {
	s.state.world.NextGen(s.state.rule)
}

func (s *myScene) Draw(drawer drawers.Drawer) {
	s.state.worldDrawer.Draw(drawer)
}

func main() {
	scene.Run(&myScene{})
}
