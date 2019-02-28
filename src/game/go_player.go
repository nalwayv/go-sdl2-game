package game

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

/*
* Implements IGameObject interface
 - Update()
 - Clean()
 - HandleInput()
 - Load()
*/

// Player ...
type Player struct {
	obj *SdlGameObject // inherit game object
}

// NewPlayer .. constructor initialise a new player object
func NewPlayer(params *LoadParams) *Player {

	player := &Player{}

	player.obj = NewSdlGObj(params) // init

	return player
}

// Load ...
// func (p *Player) Load() {
// }

// Draw ...
func (p *Player) Draw() {
	p.obj.Draw()
}

// Update ...
func (p *Player) Update() {
	p.obj.Velocity.SetX(0)
	p.obj.Velocity.SetY(0)

	// mouse to mouse pinter on screen
	// vec := SInputHandler.GetMousePosition()
	// p.obj.Velocity = vec2d.Divide(
	// 	*vec2d.Sub(*vec, *p.obj.Position), // mouse pos - player pos
	// 	100,                               // dampener  - slow down
	// )

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
		// Xvalue(0,1) -> 0 = main joystick, 1 = left analog stick
		// Yvalue(0,2) -> 0 = main joystick, 2 = right analog stick
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
		fmt.Println("player up pressed")
		p.obj.Velocity.SetY(-2)
	}

	if SInputHandler.IsKeyDown(sdl.SCANCODE_DOWN) {
		fmt.Println("player down pressed")
		p.obj.Velocity.SetY(2)
	}

	if SInputHandler.IsKeyDown(sdl.SCANCODE_LEFT) {
		fmt.Println("player left pressed")
		p.obj.Velocity.SetX(-2)
	}

	if SInputHandler.IsKeyDown(sdl.SCANCODE_RIGHT) {
		fmt.Println("player right pressed")
		p.obj.Velocity.SetX(2)
	}

}
