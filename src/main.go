package main

import (
	"./game"
	"github.com/veandco/go-sdl2/sdl"
)

//------------------------------------

const WINDOW_WIDTH int32 = 640
const WINDOW_HEIGHT int32 = 480

//------------------------------------

func run() {
	game.STheGame.Init(
		"sdl",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		WINDOW_WIDTH,
		WINDOW_HEIGHT,
		false,
	)

	for game.STheGame.Running {
		game.STheGame.HandleEvents()
		game.STheGame.Update()
		game.STheGame.Render()
	}

	game.STheGame.Clean()
}

func main() {
	run()
}

//------------------------------------
