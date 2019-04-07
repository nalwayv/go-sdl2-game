package game

/*
IGameObject interface
---

- Draw()
- Update()
- Clean()
- Load(*params)
*/

// Enemy ...
type Enemy struct {
	obj *GameObject // inherit game object
}

// NewEnemy .. constructor
func NewEnemy() *Enemy {
	enemy := &Enemy{}
	enemy.obj = NewGameObject()
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
	e.obj.Clean()
}

// HandleInput ...
func (e *Enemy) HandleInput() {

}
