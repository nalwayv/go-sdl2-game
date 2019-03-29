package game

/*
Info
---
Singleton Main game

**/

import (
	"sync"

	"../gologger"
	"github.com/veandco/go-sdl2/sdl"
)

var (
	gm     *Game
	gmOnce sync.Once
)

// STheGame ... interact with singleton
var STheGame = newGame()

// Game ...
type Game struct {
	Running      bool
	Window       *sdl.Window
	Renderer     *sdl.Renderer
	GameObject   []IGameObject
	StateMachine *StateMachine
	Width        int32
	Height       int32
}

// New ... create singleton
func newGame() *Game {
	gologger.SLogger.Println("Init NewGame")
	gmOnce.Do(func() {
		gm = &Game{}
	})
	return gm
}

// GetRenderer ...
func (g *Game) GetRenderer() *sdl.Renderer {
	return g.Renderer
}

// GetStateMachine ...
func (g *Game) GetStateMachine() *StateMachine {
	return g.StateMachine
}

// GetWidth ...
func (g *Game) GetWidth() int32 {
	return g.Width
}

// GetHeight ...
func (g *Game) GetHeight() int32 {
	return g.Height
}

// Init ...
func (g *Game) Init(title string, xPos, yPos, width, height int32, fullscreen bool) {
	var err error
	var flag uint32

	// set fullscreen or not
	if fullscreen {
		flag = sdl.WINDOW_FULLSCREEN
	} else {
		flag = sdl.WINDOW_SHOWN
	}

	// initialize sdl2
	err = sdl.Init(sdl.INIT_EVERYTHING)
	checkError(err)

	// create window
	g.Width = width
	g.Height = height
	g.Window, err = sdl.CreateWindow(title, xPos, yPos, g.Width, g.Height, flag)
	checkError(err)

	// create renderer
	g.Renderer, err = sdl.CreateRenderer(g.Window, -1, 0)
	checkError(err)

	// set bg color
	err = g.Renderer.SetDrawColor(100, 149, 237, 255) // cornflower blue
	checkError(err)

	g.Running = true

	// input singleton
	SInputHandler.InitialiseJoySticks()

	// init game factory objects
	STheGameObjFactory.Register("MenuButton", NewMenuButtonCreator())
	STheGameObjFactory.Register("Player", NewPlayerCreator())
	STheGameObjFactory.Register("Enemy", NewEnemyCreator())
	STheGameObjFactory.Register("AnimatedGraphic", NewAnimGraphicCreator())

	// state machine / set to menu state
	g.StateMachine = NewStateMachine()
	g.GetStateMachine().ChangeState(NewMenuState())
}

// Render ...
func (g *Game) Render() {
	g.Renderer.Clear()

	g.GetStateMachine().Render()

	g.Renderer.Present()
}

// Update ...
func (g *Game) Update() {
	g.GetStateMachine().Update()
}

// HandleEvents ...
func (g *Game) HandleEvents() {
	SInputHandler.Update()

	if SInputHandler.IsKeyDown(sdl.SCANCODE_RETURN) {
		g.GetStateMachine().ChangeState(NewPlayState())
	}
}

// Clean ...
func (g *Game) Clean() {
	g.Window.Destroy()
	g.Renderer.Destroy()
	SInputHandler.Clean()
	sdl.Quit()
}

// Quit ...
func (g *Game) Quit() {
	gologger.SLogger.Println("Quit App")
	g.Running = false
}
