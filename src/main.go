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

// game loop update
var (
	previousTime uint32
	lagTime      uint32
	currentTime  uint32
	elapsedTime  uint32
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

		currentTime = sdl.GetTicks()
		elapsedTime = currentTime - previousTime
		previousTime = currentTime
		lagTime += elapsedTime

		// Update only every Milliseconds per frame.
		// If lag larger then update frames, update until caught up.
		for (lagTime >= DelayTime) && game.STheGame.Running {

			game.STheGame.HandleEvents()
			game.STheGame.Update()

			lagTime -= DelayTime
		}

		game.STheGame.Render()
	}

	// clean up
	game.STheGame.Clean()
}

func main() {
	run()
}

//------------------------------------
