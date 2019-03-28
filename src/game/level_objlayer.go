package game

/*
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

// Render ...
func (ol *ObjectLayer) Render() {
	for _, v := range ol.gameObjects {
		v.Draw()
	}
}

// Update ...
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
