package game

// Level ...
type Level struct{
	tileSet []*Tileset
	layers []ILayer
}

// NewLevel ...
func NewLevel()*Level{
	l:=Level{}

	l.tileSet = make([]*Tileset, 0)
	l.layers = make([]ILayer, 0)

	return &l
}

// Update ...
func(l *Level)Update(){
	for _,v := range l.layers{
		v.Update()
	}

}

// Render ...
func(l *Level)Render(){
	for _,v := range l.layers{
		v.Render()
	}
}

// GetTileSet ...
func(l *Level)GetTileSet()[]*Tileset{
	return l.tileSet
}

// GetLayers ...
func(l *Level)GetLayers()[]ILayer{
	return l.layers
}
