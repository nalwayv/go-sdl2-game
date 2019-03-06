package main

import (
	"./game"
	"github.com/veandco/go-sdl2/sdl"
)

//------------------------------------

// Information ...
const (
	WindowWidth  int32  = 640
	WindowHeight int32  = 480
	Fps          uint32 = 60
	DelayTime    uint32 = 1000.0 / Fps
	AccTime      uint32 = 0.0
)

//------------------------------------

func run() {
	// singleton init game obj
	game.STheGame.Init(
		"sdl",                   // name
		sdl.WINDOWPOS_UNDEFINED, // window position
		sdl.WINDOWPOS_UNDEFINED, // window position
		WindowWidth,             // width
		WindowHeight,            // height
		false,                   // full screen
	)

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
