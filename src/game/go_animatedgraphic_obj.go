package game

/*
Implements IGameObject interface

- Draw()
- Update()
- Clean()
- Load()
*/

// AnimatedGraphic ...
type AnimatedGraphic struct {
	obj *SdlGameObject
}

// NewAnimatedGraphic ...
func NewAnimatedGraphic(params *LoadParams) *AnimatedGraphic {
	ag := &AnimatedGraphic{}

	ag.obj = NewSdlGObj(params)
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
	// val := (1000 / ag.obj.animSpeed) % ag.obj.NumFrames

	// ag.obj.CurrentFrame = int32(ticks / val)
}

// Clean ...
func (ag *AnimatedGraphic) Clean() {
	ag.obj.Clean()
}

// Load ...
func (ag *AnimatedGraphic) Load(param *LoadParams) {

}
