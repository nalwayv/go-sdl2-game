package game

// Tileset ...
type Tileset struct {
	GridID     int
	TileWidth  int
	TileHeight int
	Spacing    int
	Margin     int
	Width      int
	Height     int
	NumColumns int
	Name       string
}

// NewTileset ...
func NewTileset() *Tileset {
	return &Tileset{}
}
