package main

import (
	"./game"
	"github.com/veandco/go-sdl2/sdl"
)

//------------------------------------

const (
	WINDOW_WIDTH  int32 = 640
	WINDOW_HEIGHT int32 = 480
)

//------------------------------------

func run() {
	ggame := &game.Game{}
	defer ggame.Clean()

	ggame.Init("sdl",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		WINDOW_WIDTH,
		WINDOW_HEIGHT,
		false)

	for ggame.Running {
		ggame.HandleEvents()
		ggame.Update()
		ggame.Render()
	}
}

func main() {
	run()
}

//------------------------------------
