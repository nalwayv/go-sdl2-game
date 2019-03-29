package game

/*
* Info
* ---
* store game objects used by the current state and
* calls their Update and Draw functions from its Render and Update
*
* ILayer interface
* ---
* - Render()
* - Update()
* */

// ObjectLayer ...
type ObjectLayer struct {
	gameObjects []IGameObject
}

// NewObjectLayer ...
func NewObjectLayer() *ObjectLayer {
	ol := ObjectLayer{}

	ol.gameObjects = make([]IGameObject, 0)

	return &ol
}

// Render ... call each obj's Draw
func (ol *ObjectLayer) Render() {
	for _, v := range ol.gameObjects {
		v.Draw()
	}
}

// Update ... call each obj's Update
func (ol *ObjectLayer) Update() {
	for _, v := range ol.gameObjects {
		v.Update()
	}
}

// GetGameObjs ...
func (ol *ObjectLayer) GetGameObjs() []IGameObject {
	return ol.gameObjects
}

// PushOntoGameObj ...
func (ol *ObjectLayer) PushOntoGameObj(value IGameObject) {
	ol.gameObjects = append(ol.gameObjects, value)
}
