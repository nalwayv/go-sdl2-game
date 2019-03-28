package game

/*
* ILayer interface
* ---
* - Render()
* - Update()
**/

import (
	"../gologger"
	"../vec2d"
)

// TileLayer ...
type TileLayer struct {
	tileSize   int             // size of a tile
	numColumns int             // number of columns on screen
	numRows    int             // number of rows on screen
	position   *vec2d.Vector2D // position on the tile map to start drawing from
	velocity   *vec2d.Vector2D // velocity of tileset on screen

	tileSets []*Tileset // data from parsed json file
	tileIDs  [][]int    // data from parsed json file
}

// NewTileLayer ...
func NewTileLayer(tileSize int, tileset []*Tileset) *TileLayer {
	tl := TileLayer{}

	tl.tileSize = tileSize
	tl.tileSets = tileset
	tl.position = vec2d.NewVector2d(0, 0) // top|left
	tl.velocity = vec2d.NewVector2d(0, 0)

	// number of columns and rows currently displayed on screen
	tl.numColumns = (int(STheGame.GetWidth()) / tl.tileSize)
	tl.numRows = (int(STheGame.GetHeight()) / tl.tileSize)

	return &tl
}

// Render ... ILayer interface
func (tl *TileLayer) Render() {

	// float64 -> int
	x1 := int(tl.position.GetX()) / tl.tileSize
	y1 := int(tl.position.GetY()) / tl.tileSize

	x2 := int(tl.position.GetX()) % tl.tileSize
	y2 := int(tl.position.GetY()) % tl.tileSize

	for r := 0; r < tl.numRows; r++ {
		for c := 0; c < tl.numColumns; c++ {
			// tile index number from tiled
			id := tl.tileIDs[r+y1][c+x1]

			// nothing to draw
			if id == 0 {
				continue
			}

			tileset := tl.GetTilesetByID(id)
			id-- // tiled adds one to tile id so negative one to get the correct tile

			STextureManager.DrawTile(
				tileset.Name,
				int32(tileset.Margin),     // margin around each tile
				int32(tileset.Spacing),    // spacing between each tile
				int32((c*tl.tileSize)-x2), // x
				int32((r*tl.tileSize)-y2), // y
				int32(tl.tileSize),        // width
				int32(tl.tileSize),        // height
				int32((id-(tileset.GridID-1))/tileset.NumColumns), // location of the tile on the tilesheet x,y
				int32((id-(tileset.GridID-1))%tileset.NumColumns),
				STheGame.GetRenderer()) // renderer
		}
	}
}

//Update ... ILayer interface
func (tl *TileLayer) Update() {
	tl.position = vec2d.Add(*tl.position, *tl.velocity)
	//tl.velocity.SetX(1)
}

// SetTileIDs ... set the ids of tiles
func (tl *TileLayer) SetTileIDs(data [][]int) {
	tl.tileIDs = data
}

// SetTileSize ... set the size of a tile
func (tl *TileLayer) SetTileSize(value int) {
	tl.tileSize = value
}

// GetTilesetByID ... return tileset from list of tilesets by id
func (tl *TileLayer) GetTilesetByID(tileID int) *Tileset {
	n := len(tl.tileSets)

	for i := 0; i < n; i++ {
		// if in bounds
		if i+1 <= n-1 {

			current := tl.tileSets[i].GridID
			next := tl.tileSets[i+1].GridID

			// if in order
			if tileID >= current && tileID < next {
				return tl.tileSets[i]
			}

		} else {
			return tl.tileSets[i]
		}
	}

	// blank
	gologger.SLogger.Println("created blank tileset")
	t := Tileset{}
	return &t
}
