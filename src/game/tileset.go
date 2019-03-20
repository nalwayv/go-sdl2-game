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
func NewTileset(gid, width, height, tilewidth, tileheight, spacing, margin, numcolumns int, name string) *Tileset {
	return &Tileset{
		GridID:     gid,
		TileWidth:  tilewidth,
		TileHeight: tileheight,
		Width:      width,
		Height:     height,
		Spacing:    spacing,
		Margin:     margin,
		NumColumns: numcolumns,
		Name:       name,
	}
}
