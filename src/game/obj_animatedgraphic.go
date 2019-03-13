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
func NewAnimatedGraphic() *AnimatedGraphic {
	ag := &AnimatedGraphic{}
	ag.obj = NewSdlGObj()
	return ag
}
// Load ...
func (ag *AnimatedGraphic) Load(param *LoadParams) {

}

// Draw ...
func (ag *AnimatedGraphic) Draw() {
	ag.obj.Draw()
}

// Update ...
func (ag *AnimatedGraphic) Update() {
	// TODO :: zero error

	// ticks := sdl.GetTicks()
	// val := (1000 / ag.obj.animSpeed) % ag.obj.NumFrames

	// ag.obj.CurrentFrame = int32(ticks / val)
}

// Clean ...
func (ag *AnimatedGraphic) Clean() {
	ag.obj.Clean()
}

