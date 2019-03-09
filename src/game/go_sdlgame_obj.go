package game

/*
Implements IGameObject interface

- Draw()
- Update()
- Clean()
*/


import (
	"./vec2d"
	"github.com/veandco/go-sdl2/sdl"
)

// SdlGameObject ...
type SdlGameObject struct {
	//*GObject

	ID           string
	Width        int32
	Height       int32
	CurrentFrame int32
	CurrentRow   int32
	NumFrames    int

	Position     *vec2d.Vector2D
	Velocity     *vec2d.Vector2D
	Acceleration *vec2d.Vector2D
}

// NewSdlGObj ...
func NewSdlGObj(params *LoadParams) *SdlGameObject {
	obj := &SdlGameObject{}

	obj.Position = vec2d.NewVector2d(float64(params.X()), float64(params.Y()))
	obj.Velocity = vec2d.NewVector2d(0.0, 0.0)
	obj.Acceleration = vec2d.NewVector2d(0.0, 0.0)

	obj.Width = params.Width()
	obj.Height = params.Height()

	obj.ID = params.ID()

	obj.CurrentRow = 1
	obj.CurrentFrame = 1
	obj.NumFrames = params.NumFrames()

	return obj
}

// Draw ...
func (g *SdlGameObject) Draw() {
	// flipped or not based on velocity
	if g.Velocity.GetX() > 0 {
		STextureManager.DrawFrame(
			g.ID,
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
			g.ID,
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
func (g *SdlGameObject) Update() {
	g.Velocity = vec2d.Add(*g.Velocity, *g.Acceleration)
	g.Position = vec2d.Add(*g.Position, *g.Velocity)
}

// Clean ...
func (g *SdlGameObject) Clean() {
}
