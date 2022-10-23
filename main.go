package main

import (
	"github.com/zronev/cellular-zoo/renderers"
	"github.com/zronev/cellular-zoo/loop"

	"github.com/faiface/pixel/pixelgl"
)

func run() {
	wr := new(renderers.WindowRenderer)
	err := wr.Setup()
	if err != nil {
		panic(err)
	}
	loop.Loop(wr)
}


func main() {
	pixelgl.Run(run)
}
