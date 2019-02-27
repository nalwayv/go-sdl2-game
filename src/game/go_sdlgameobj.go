package game

import "../vec2d"

// SdlGameObject ...
type SdlGameObject struct {
	//*GObject

	ID           string
	Width        int32
	Height       int32
	CurrentFrame int32
	CurrentRow   int32

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

	return obj
}

// Draw ...
func (g *SdlGameObject) Draw() {
	STextureManager.DrawFrame(
		g.ID,                     // texture id
		int32(g.Position.GetX()), // xpos
		int32(g.Position.GetY()), // ypos
		g.Width,                  // width of img
		g.Height,                 // height of img
		g.CurrentRow,             // current frame row
		g.CurrentFrame,           // current frame
		STheGame.GetRenderer(),   // renderer
		0,                        // flipped
	)
}

// Update ...
func (g *SdlGameObject) Update() {
	g.Velocity = vec2d.Add(*g.Velocity, *g.Acceleration)
	g.Position = vec2d.Add(*g.Position, *g.Velocity)
}

// Clean ...
func (g *SdlGameObject) Clean() {
}
