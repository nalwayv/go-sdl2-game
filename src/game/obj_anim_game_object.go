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
	"../vec2d"
	"github.com/veandco/go-sdl2/sdl"
)

// AnimObject ...
type AnimObject struct {
	GameObject        // inherit GameObject
	frames     *Frame // animation
}

// NewAnimObject ...
func NewAnimObject() *AnimObject {
	obj := &AnimObject{}

	return obj
}

// Load ... set up variables
func (g *AnimObject) Load(params *LoadParams) {

	// --- same as GameObject
	g.Position = vec2d.NewVector2d(float64(params.GetX()), float64(params.GetY()))
	g.Velocity = vec2d.NewVector2d(0.0, 0.0)
	g.Acceleration = vec2d.NewVector2d(0.0, 0.0)

	g.Width = params.GetWidth()
	g.Height = params.GetHeight()

	g.TextureID = params.GetTextureID()

	g.CurrentRow = 1
	g.CurrentFrame = 1
	g.NumFrames = params.GetNumFrames()
	g.AnimSpeed = params.GetAnimationSpeed()

	// --- anim

	g.frames = NewFrame()
	g.frames.SetFrame(
		g.CurrentFrame,
		g.CurrentRow,
		g.CurrentFrame,
		g.CurrentRow+3,
		g.AnimSpeed,
		g.NumFrames)
}

// Draw ...
func (g *AnimObject) Draw() {
	// flipped or not based on velocity
	if g.Velocity.GetX() > 0 {
		g.frames.DrawFrame(
			g.TextureID,
			int32(g.Position.GetX()),
			int32(g.Position.GetY()),
			g.Width,
			g.Height,
			STheGame.GetRenderer(),
			sdl.FLIP_HORIZONTAL)
	} else {
		g.frames.DrawFrame(
			g.TextureID,
			int32(g.Position.GetX()),
			int32(g.Position.GetY()),
			g.Width,
			g.Height,
			STheGame.GetRenderer(),
			0)
	}
}

// Update ...
func (g *AnimObject) Update() {
	g.Velocity = vec2d.Add(*g.Velocity, *g.Acceleration)
	g.Position = vec2d.Add(*g.Position, *g.Velocity)

	// update animation
	g.frames.UpdateFrame()
}

// Clean ...
func (g *AnimObject) Clean() {
}
