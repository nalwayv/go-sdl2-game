package game

import (
	"fmt"
	"sync"

	"../gologger"
	"github.com/veandco/go-sdl2/sdl"
)

// Game ...
type Game struct {
	Running  bool
	Window   *sdl.Window
	Renderer *sdl.Renderer
	// P1           *Player
	// E1           *Enemy
	GameObject   []IGameObject
	StateMachine *StateMachine
}

var gm *Game
var gOnce sync.Once

// STheGame ... interact with singleton
var STheGame = newGame()

// New ...
func newGame() *Game {
	gOnce.Do(func() {
		gm = &Game{}
	})
	return gm
}

// GetRenderer ...
func (g *Game) GetRenderer() *sdl.Renderer {
	return g.Renderer
}

// Init ...
func (g *Game) Init(title string, xPos, yPos, width, height int32, fullscreen bool) {
	/*
		FLAGS INIT
		----------
		SDL_INIT_HAPTIC 			Force feedback subsystem
		SDL_INIT_AUDIO 				Audio subsystem
		SDL_INIT_VIDEO 				Video subsystem
		SDL_INIT_TIMER 				Timer subsystem
		SDL_INIT_JOYSTICK 			Joystick subsystem
		SDL_INIT_EVERYTHING 		All subsystems
		SDL_INIT_NOPARACHUTE 		Don't catch fatal signals

		FLAGS RENDERER
		--------------
		SDL_RENDERER_SOFTWARE 		Use software rendering
		SDL_RENDERER_ACCELERATED 	Use hardware acceleration
		SDL_RENDERER_PRESENTVSYNC 	Synchronize renderer update with screen's refresh rate
		SDL_RENDERER_TARGETTEXTURE	Supports render to texture

		FLAGS WINDOW
		-------------
		SDL_WINDOW_FULLSCREEN 		Make the window fullscreen
		SDL_WINDOW_OPENGL 			Window can be used with as an OpenGL context
		SDL_WINDOW_SHOWN 			The window is visible
		SDL_WINDOW_HIDDEN 			Hide the window
		SDL_WINDOW_BORDERLESS 		No border on the window
		SDL_WINDOW_RESIZABLE 		Enable resizing of the window
		SDL_WINDOW_MINIMIZED 		Minimize the window
		SDL_WINDOW_MAXIMIZED 		Maximize the window
		SDL_WINDOW_INPUT_GRABBED 	Window has grabbed input focus
		SDL_WINDOW_INPUT_FOCUS 		Window has input focus
		SDL_WINDOW_MOUSE_FOCUS 		Window has mouse focus
	*/

	var err error
	var flag uint32

	if fullscreen {
		flag = sdl.WINDOW_FULLSCREEN
	} else {
		flag = sdl.WINDOW_SHOWN
	}

	err = sdl.Init(sdl.INIT_EVERYTHING)
	checkError(err)

	g.Window, err = sdl.CreateWindow(title, xPos, yPos, width, height, flag)
	checkError(err)

	g.Renderer, err = sdl.CreateRenderer(g.Window, -1, 0)
	checkError(err)

	err = g.Renderer.SetDrawColor(255, 255, 255, 255)
	checkError(err)

	g.Running = true

	// texture singleton
	STextureManager.Load("assets/handheld.png", "animate", g.Renderer)

	// input singleton
	SInputHandler.InitialiseJoySticks()

	// // player
	// g.P1 = NewPlayer(NewParams(0, 0, 96, 96, "animate"))

	// // enemy
	// g.E1 = NewEnemy(NewParams(100, 100, 96, 96, "animate"))

	// // add gameobjects
	// g.GameObject = make([]IGameObject, 2)
	// g.GameObject = append(g.GameObject, g.P1)
	// g.GameObject = append(g.GameObject, g.E1)

	// statemachine
	g.StateMachine = NewStateMachine()
	g.StateMachine.ChangeState(NewMenuState())
}

// Render ...
func (g *Game) Render() {
	g.Renderer.Clear()

	// for _, v := range g.Gobjs {
	// 	switch v.(type) {
	// 	case *Player:
	// 		v.(*Player).Draw()

	// 	case *Enemy:
	// 		v.(*Enemy).Draw()
	// 	}
	// }

	g.StateMachine.Render()

	g.Renderer.Present()
}

// Update ...
func (g *Game) Update() {
	// for _, v := range g.Gobjs {
	// 	switch v.(type) {

	// 	case *Player:
	// 		v.(*Player).Update()

	// 	case *Enemy:
	// 		v.(*Enemy).Update()
	// 	}
	// }

	g.StateMachine.Update()
}

// HandleEvents ...
func (g *Game) HandleEvents() {
	SInputHandler.Update()

	if SInputHandler.IsKeyDown(sdl.SCANCODE_RETURN) {
		g.StateMachine.ChangeState(NewPlayState())
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
	fmt.Println("Quit")
	gologger.SLogger.Println("Quit App")
	g.Running = false
}
