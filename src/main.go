package main

import (
	"./game"
	"github.com/veandco/go-sdl2/sdl"
)

//------------------------------------

// Settings ...
const (
	WindowWidth  int32  = 640
	WindowHeight int32  = 480
	Fps          uint32 = 60
	DelayTime    uint32 = 1000.0 / Fps
)

//------------------------------------

func run() {
	// singleton init game obj
	game.STheGame.Init(
		"sdl",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		WindowWidth,
		WindowHeight,
		false)

	// main loop
	for game.STheGame.Running {

		framesStart := sdl.GetTicks()

		game.STheGame.HandleEvents()
		game.STheGame.Update()
		game.STheGame.Render()

		framesEnd := sdl.GetTicks() - framesStart

		if framesEnd < DelayTime {
			sdl.Delay(DelayTime - framesEnd)
		}
	}

	// clean up
	game.STheGame.Clean()
}

func main() {
	run()
}

//------------------------------------
