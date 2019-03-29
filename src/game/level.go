package game

/*
* Info
* ---
* Used to store data on current level
*
* layers[] :: store current levels game objects and tile layers
* from objLayer and tilelayer that store info on current levels game objects and tiles
*
* Interface ILayer
* ---
* - Update()
* - Render()
* */

// Level ...
type Level struct {
	tileSets []*Tileset // tilesets for this level
	layers   []ILayer   // objects that implement ILayer interface
}

// NewLevel ...
func NewLevel() *Level {
	l := Level{}

	l.tileSets = make([]*Tileset, 0)
	l.layers = make([]ILayer, 0)

	return &l
}

// Update ... call any object within layers Update
func (l *Level) Update() {
	for _, v := range l.layers {
		v.Update()
	}
}

// Render ... call any object within layers Render
func (l *Level) Render() {
	for _, v := range l.layers {
		v.Render()
	}
}

// SetTileSet ... populate tileset slice
func (l *Level) SetTileSet(ts []*Tileset) {
	l.tileSets = ts
}

// GetTileSet ...
func (l *Level) GetTileSet() []*Tileset {
	return l.tileSets
}

// --- get/set

// AppendToTileSet ... added so i dont have to use a pointer to this slice
func (l *Level) AppendToTileSet(ts *Tileset) {
	l.tileSets = append(l.tileSets, ts)
}

// SetLayers ...
func (l *Level) SetLayers(val []ILayer) {
	l.layers = val
}

// GetLayers ...
func (l *Level) GetLayers() []ILayer {
	return l.layers
}

// AppendToLayer ... added so i dont have to use a pointer to this slice
func (l *Level) AppendToLayer(tl ILayer) {
	l.layers = append(l.layers, tl)
}
