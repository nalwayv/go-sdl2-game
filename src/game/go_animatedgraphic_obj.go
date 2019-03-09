package game

/*
Implements IGameObject interface

- Draw()
- Update()
- Clean()
*/

// AnimatedGraphic ...
type AnimatedGraphic struct {
	obj            *SdlGameObject
	animationSpeed int
}

// NewAnimatedGraphic ...
func NewAnimatedGraphic(params *LoadParams, speed int) *AnimatedGraphic {
	ag := &AnimatedGraphic{}

	ag.obj = NewSdlGObj(params)
	ag.animationSpeed = speed
	return ag
}

// Draw ...
func (ag *AnimatedGraphic) Draw() {
	ag.obj.Draw()
}

// Update ...
func (ag *AnimatedGraphic) Update() {
	// TODO - zero error

	// ticks := sdl.GetTicks()
	// val := (1000 / ag.animationSpeed) % ag.obj.NumFrames

	// ag.obj.CurrentFrame = int32(ticks / val)
}

// Clean ...
func (ag *AnimatedGraphic) Clean() {
	ag.obj.Clean()
}
