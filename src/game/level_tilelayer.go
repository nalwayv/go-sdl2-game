package game

/*
*ILayer interface
---
- Render()
- Update()
* */

import (
	"../gologger"
	"./vec2d"
)

// TileLayer ...
type TileLayer struct {
	tileSize   int
	numColumns int
	numRows    int
	position   *vec2d.Vector2D
	velocity   *vec2d.Vector2D

	tileSets []*Tileset // data from parsed json file
	tileIDs  [][]int    // data from parsed json file
}

// NewTileLayer ...
func NewTileLayer(tileSize int, tileset []*Tileset) *TileLayer {
	tl := TileLayer{}

	tl.tileSize = tileSize
	tl.tileSets = tileset
	tl.position = vec2d.NewVector2d(0, 0)
	tl.velocity = vec2d.NewVector2d(0, 0)

	// set to size of screen so only to draw what is on screen not off it
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
			id := tl.tileIDs[r+y1][c+x1]

			// nothing to draw
			if id == 0 {
				continue
			}

			tileset := tl.GetTilesetByID(id)
			id--

			// TODO - draw tiles
			STextureManager.DrawTile(
				tileset.Name,
				int32(tileset.Margin),
				int32(tileset.Spacing),
				int32((c*tl.tileSize)-x2), // x position
				int32((r*tl.tileSize)-y2), // y position
				int32(tl.tileSize),        // width
				int32(tl.tileSize),        // height
				int32((id-(tileset.GridID-1))/tileset.NumColumns), // current row
				int32((id-(tileset.GridID-1))%tileset.NumColumns), // current frame
				STheGame.GetRenderer()) // renderer

		}
	}
}

//Update ... ILayer interface
func (tl *TileLayer) Update() {
	tl.position = vec2d.Add(*tl.position, *tl.velocity)
}

// SetTileIDs ... set the ids of tiles
func (tl *TileLayer) SetTileIDs(data [][]int) {
	tl.tileIDs = data
}

// SetTileSize ... set the size of a tile
func (tl *TileLayer) SetTileSize(value int) {
	tl.tileSize = value
}

// GetTilesetByID ... return tileset ffrom list of tilesets by id
func (tl *TileLayer) GetTilesetByID(tileID int) *Tileset {
	n := len(tl.tileSets)

	// if in bounds
	for i := 0; i < n; i++ {

		if i+1 <= n-1 {
			// if tileID is the same as current gridID,
			// andis also a lower value then the next gridID

			current := tl.tileSets[i].GridID
			next := tl.tileSets[i+1].GridID

			if tileID >= current && tileID < next {
				return tl.tileSets[i]
			}
		} else {
			return tl.tileSets[i]
		}

	}

	// empty
	gologger.SLogger.Println("did not find tileset retruning empty tileset.")
	t := &Tileset{}
	return t
}
