package game

/*
IGameObject interface
---

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
	ag.obj.Load(param)
}

// Draw ...
func (ag *AnimatedGraphic) Draw() {
	ag.obj.Draw()
}

// Update ...
func (ag *AnimatedGraphic) Update() {

}

// Clean ...
func (ag *AnimatedGraphic) Clean() {
	ag.obj.Clean()
}
