package game

/*
Implements IGameObject interface

- Draw()
- Update()
- Clean()
- Load()
*/

// Enemy ...
type Enemy struct {
	obj *SdlGameObject // inherit game object
}

// NewEnemy .. constructor
func NewEnemy() *Enemy {
	enemy := &Enemy{}
	enemy.obj = NewSdlGObj()
	return enemy
}

// Load ...
func (e *Enemy) Load(params *LoadParams) {
	e.obj.Load(params)
}

// Draw ...
func (e *Enemy) Draw() {
	e.obj.Draw()
}

// Update ...
func (e *Enemy) Update() {
	e.obj.Update()
}

// Clean ...
func (e *Enemy) Clean() {
}

// HandleInput ...
func (e *Enemy) HandleInput() {

}
