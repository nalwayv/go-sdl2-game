package game

/*
IGameObject interface
---

- Draw()
- Update()
- Clean()
- Load(*params)
*/

import (
	"./vec2d"
	"github.com/veandco/go-sdl2/sdl"
)

// Player ...
type Player struct {
	obj *SdlGameObject // inherit game object
}

// NewPlayer .. constructor initialise a new player object
func NewPlayer() *Player {
	player := &Player{}
	player.obj = NewSdlGObj()
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
	p.obj.Clean()
}

// HandleInput ...
func (p *Player) HandleInput() {

	// Joystick
	if SInputHandler.JoySticksInitialised() {
		// left stick left / right
		if SInputHandler.GetXvalue(0, 1) > 0 || SInputHandler.GetXvalue(0, 1) < 0 {
			vel := float64(1 * SInputHandler.GetXvalue(0, 1))
			p.obj.Velocity.SetX(vel)
		}

		// left stick up / down
		if SInputHandler.GetYvalue(0, 1) > 0 || SInputHandler.GetYvalue(0, 1) < 0 {
			vel := float64(1 * SInputHandler.GetYvalue(0, 1))
			p.obj.Velocity.SetY(vel)
		}

		// right stick left / right
		if SInputHandler.GetXvalue(0, 2) > 0 || SInputHandler.GetXvalue(0, 2) < 0 {
			vel := float64(1 * SInputHandler.GetXvalue(0, 2))
			p.obj.Velocity.SetX(vel)
		}

		// right stick up / down
		if SInputHandler.GetYvalue(0, 2) > 0 || SInputHandler.GetYvalue(0, 2) < 0 {
			vel := float64(1 * SInputHandler.GetYvalue(0, 2))
			p.obj.Velocity.SetY(vel)
		}

		// button
		// joystick 1, Y button
		if SInputHandler.GetButtonState(0, 3) {
			p.obj.Velocity.SetX(1)
		}
	}

	// No joystick

	// mouse left
	if SInputHandler.GetMouseButtonState(0) {
		p.obj.Velocity.SetX(1)
	}

	// keyboard
	if SInputHandler.IsKeyDown(sdl.SCANCODE_UP) {
		p.obj.Velocity.SetY(-2)
	}

	if SInputHandler.IsKeyDown(sdl.SCANCODE_DOWN) {
		p.obj.Velocity.SetY(2)
	}

	if SInputHandler.IsKeyDown(sdl.SCANCODE_LEFT) {
		p.obj.Velocity.SetX(-2)
	}

	if SInputHandler.IsKeyDown(sdl.SCANCODE_RIGHT) {
		p.obj.Velocity.SetX(2)
	}

	// follow mouse
	target := SInputHandler.GetMousePosition()
	p.obj.Velocity = vec2d.Divide(*vec2d.Sub(*target, *p.obj.Position), 100)

}
