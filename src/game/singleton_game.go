package game

/*
* Singleton
* ---
* Main game initialization  
**/

import (
	"sync"

	"../gologger"
	"github.com/veandco/go-sdl2/sdl"
)

// Game ...
type Game struct {
	Running      bool
	Window       *sdl.Window
	Renderer     *sdl.Renderer
	GameObject   []IGameObject
	StateMachine *StateMachine
	Width int32
	Height int32
}

var (
	gm    *Game
	gOnce sync.Once
)

// STheGame ... interact with singleton
var STheGame = newGame()

// New ... create singleton
func newGame() *Game {
	gologger.SLogger.Println("Init NewGame")
	gOnce.Do(func() {
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
func (g *Game)GetWidth()int32{
	return g.Width
}

// GetHeight ...
func (g *Game)GetHeight()int32{
	return g.Height
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
	err = g.Renderer.SetDrawColor(255, 255, 255, 255)
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
