package game

// Level ...
type Level struct {
	tileSets []*Tileset
	layers   []ILayer
}

// NewLevel ...
func NewLevel() *Level {
	l := Level{}

	l.tileSets = make([]*Tileset, 0)
	l.layers = make([]ILayer, 0)

	return &l
}

// Update ...
func (l *Level) Update() {
	for _, v := range l.layers {
		v.Update()
	}
}

// Render ...
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

// AppendToTileLayer ... added so i dont have to use a pointer to this slice
func (l *Level) AppendToTileLayer(tl ILayer) {
	l.layers = append(l.layers, tl)
}
