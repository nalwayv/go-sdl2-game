package main

import (
	"./game"
	"./gologger"
	"github.com/veandco/go-sdl2/sdl"
)

//------------------------------------

const (
	WINDOW_WIDTH  int32 = 640
	WINDOW_HEIGHT int32 = 480
)

// Globals ...
var (
	golog = gologger.GetInstance("src/gologger/golog.log")
)

//------------------------------------

// checkError ... check for errors
func checkError(err error) {
	if err != nil {
		golog.Fatalln(err)
	}
}

func run() {
	var err error
	ggame := &game.Game{}
	defer ggame.Clean()

	err = ggame.Init("sdl",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		WINDOW_WIDTH,
		WINDOW_HEIGHT,
		false)
	checkError(err)

	for ggame.Running {
		ggame.HandleEvents()
		ggame.Update()
		ggame.Render()
	}
}

//------------------------------------
func main() {
	run()
}

//------------------------------------
