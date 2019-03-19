package game

/*
*ILayer interface
---
- Render()
- Update()
* */

import "./vec2d"

// TileLayer ...
type TileLayer struct{
	tileSize int
	numColumns int
	numRows int
	position *vec2d.Vector2D
	velocity *vec2d.Vector2D
	tileSets []*Tileset
	tileIDs	[][]int
}

// NewTileLayer ...
func NewTileLayer(tileSize int, tileset []*Tileset)*TileLayer{
	tl := TileLayer{}

	tl.tileSize = tileSize
	tl.tileSets = tileset
	tl.position = vec2d.NewVector2d(0,0)
	tl.velocity = vec2d.NewVector2d(0,0)

	return &tl
}

// Render ...
func (tl *TileLayer) Render(){}

//Update ...
func (tl *TileLayer) Update(){}

// SetTileIDs ... set the ids of tiles
func (tl *TileLayer) SetTileIDs(data [][]int){
	tl.tileIDs = data
}

// SetTileSize ... set the size of a tile
func (tl *TileLayer)SetTileSize(value int){
	tl.tileSize = value
}

// GetTilesetByID ... return tileset ffrom list of tilesets by id
func (tl *TileLayer) GetTilesetByID(tileID int)*Tileset{
	return tl.tileSets[tileID]
}
