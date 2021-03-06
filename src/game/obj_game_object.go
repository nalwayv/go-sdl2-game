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

// GameObject ...
type GameObject struct {
	TextureID    string
	Width        int32
	Height       int32
	CurrentFrame int32
	CurrentRow   int32
	NumFrames    int
	AnimSpeed    int

	Position     *vec2d.Vector2D
	Velocity     *vec2d.Vector2D
	Acceleration *vec2d.Vector2D
}

// NewGameObject ...
func NewGameObject() *GameObject {
	obj := &GameObject{}
	return obj
}

// Load ... set up variables
func (g *GameObject) Load(params *LoadParams) {

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
}

// Draw ...
func (g *GameObject) Draw() {
	// flipped or not based on velocity
	if g.Velocity.GetX() > 0 {
		STextureManager.DrawFrame(
			g.TextureID,
			int32(g.Position.GetX()),
			int32(g.Position.GetY()),
			g.Width,
			g.Height,
			g.CurrentRow,
			g.CurrentFrame,
			STheGame.GetRenderer(),
			sdl.FLIP_HORIZONTAL,
		)
	} else {
		STextureManager.DrawFrame(
			g.TextureID,
			int32(g.Position.GetX()),
			int32(g.Position.GetY()),
			g.Width,
			g.Height,
			g.CurrentRow,
			g.CurrentFrame,
			STheGame.GetRenderer(),
			0,
		)
	}
}

// Update ...
func (g *GameObject) Update() {
	g.Velocity = vec2d.Add(*g.Velocity, *g.Acceleration)
	g.Position = vec2d.Add(*g.Position, *g.Velocity)
}

// Clean ...
func (g *GameObject) Clean() {
}
