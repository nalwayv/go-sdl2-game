package game

/*
Implements IGameObject interface

- Draw()
- Update()
- Clean()
*/

import (
	"fmt"

	"./vec2d"
	"github.com/veandco/go-sdl2/sdl"
)

// Player ...
type Player struct {
	obj *SdlGameObject // inherit game object
}

// NewPlayer .. constructor initialise a new player object
func NewPlayer(params *LoadParams) *Player {

	player := &Player{}

	//player.obj = NewSdlGObj(params) // init

	return player
}

// Load ...
func (p *Player) Load(params *LoadParams) {
	p.obj.Load(params)
}

// Draw ...
func (p *Player) Draw() {
	p.obj.Draw()
}

// Update ...
func (p *Player) Update() {
	p.obj.Velocity.SetX(0)
	p.obj.Velocity.SetY(0)

	p.HandleInput()

	p.obj.Update()
}

// Clean ...
func (p *Player) Clean() {
}

// HandleInput ...
func (p *Player) HandleInput() {

	// if joystick was found
	// else use keyboard
	if SInputHandler.JoySticksInitialised() {
		// xValue(0,1) -> 0 = main joystick, 1 = left analog stick
		// yValue(0,2) -> 0 = main joystick, 2 = right analog stick
		// left stick left / right
		if SInputHandler.Xvalue(0, 1) > 0 || SInputHandler.Xvalue(0, 1) < 0 {
			vel := float64(1 * SInputHandler.Xvalue(0, 1))
			p.obj.Velocity.SetX(vel)
		}

		// left stick up / down
		if SInputHandler.Yvalue(0, 1) > 0 || SInputHandler.Yvalue(0, 1) < 0 {
			vel := float64(1 * SInputHandler.Yvalue(0, 1))
			p.obj.Velocity.SetY(vel)
		}

		// right stick left / right
		if SInputHandler.Xvalue(0, 2) > 0 || SInputHandler.Xvalue(0, 2) < 0 {
			vel := float64(1 * SInputHandler.Xvalue(0, 2))
			p.obj.Velocity.SetX(vel)
		}

		// right stick up / down
		if SInputHandler.Yvalue(0, 2) > 0 || SInputHandler.Yvalue(0, 2) < 0 {
			vel := float64(1 * SInputHandler.Yvalue(0, 2))
			p.obj.Velocity.SetY(vel)
		}

		// button
		// joystick 1, Y button
		if SInputHandler.GetButtonState(0, 3) {
			p.obj.Velocity.SetX(1)
		}
	}

	// mouse left
	if SInputHandler.GetMouseButtonState(0) {
		p.obj.Velocity.SetX(1)
	}

	// keyboard
	if SInputHandler.IsKeyDown(sdl.SCANCODE_UP) {
		fmt.Println("player UP pressed")
		p.obj.Velocity.SetY(-2)
	}

	if SInputHandler.IsKeyDown(sdl.SCANCODE_DOWN) {
		fmt.Println("player DOWN pressed")
		p.obj.Velocity.SetY(2)
	}

	if SInputHandler.IsKeyDown(sdl.SCANCODE_LEFT) {
		fmt.Println("player LEFT pressed")
		p.obj.Velocity.SetX(-2)
	}

	if SInputHandler.IsKeyDown(sdl.SCANCODE_RIGHT) {
		fmt.Println("player RIGHT pressed")
		p.obj.Velocity.SetX(2)
	}

	// follow mouse
	target := SInputHandler.GetMousePosition()
	p.obj.Velocity = vec2d.Divide(*vec2d.Sub(*target, *p.obj.Position), 100)

}
