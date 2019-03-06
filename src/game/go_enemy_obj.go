package game

/*
	*IGameObject*
	- Draw()
 	- Update()
 	- Clean()
*/

// Enemy ...
type Enemy struct {
	obj *SdlGameObject // inherit game object
}

// NewEnemy .. constructor
func NewEnemy(params *LoadParams) *Enemy {
	enemy := &Enemy{}

	enemy.obj = NewSdlGObj(params)

	return enemy
}

// Load ...
// func (e *Enemy) Load() {
// }

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
