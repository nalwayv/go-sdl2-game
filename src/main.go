package main

import (
	"./game"
	"github.com/veandco/go-sdl2/sdl"
)

//------------------------------------

const WINDOW_WIDTH int32 = 640
const WINDOW_HEIGHT int32 = 480
const FPS uint32 = 60
const DELAY_TIME uint32 = 1000.0 / FPS

//------------------------------------

func run() {
	// singleton init game obj
	game.STheGame.Init(
		"sdl",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		WINDOW_WIDTH,
		WINDOW_HEIGHT,
		false,
	)

	for game.STheGame.Running {

		framesStart := sdl.GetTicks()
		// -
		game.STheGame.HandleEvents()
		game.STheGame.Update()
		game.STheGame.Render()
		// -
		framesEnd := sdl.GetTicks() - framesStart

		if framesEnd < DELAY_TIME {
			sdl.Delay(DELAY_TIME - framesEnd)
		}
	}

	game.STheGame.Clean()
}

func main() {
	run()
}

//------------------------------------
