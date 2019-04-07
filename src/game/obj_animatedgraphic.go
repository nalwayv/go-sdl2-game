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
	obj *GameObject
}

// NewAnimatedGraphic ...
func NewAnimatedGraphic() *AnimatedGraphic {
	ag := &AnimatedGraphic{}
	ag.obj = NewGameObject()
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
	//
}

// Clean ...
func (ag *AnimatedGraphic) Clean() {
	ag.obj.Clean()
}
