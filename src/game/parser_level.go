package game

/*
* Info
* ---
* Data parsed from the json file generated from 'Tiled'
* then passed onto the current level
**/

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"../gologger"
)

// --- JSON DATA

// JSONMap ...
type JSONMap struct {
	Tileheight int           `json:"tileheight"`
	Tilewidth  int           `json:"tilewidth"`
	Width      int           `json:"width"`
	Height     int           `json:"height"`
	Type       string        `json:"type"`
	Layers     []JSONLayers  `json:"layers"`
	Set        []JSONTileset `json:"tilesets"`
}

// JSONLayers ...
type JSONLayers struct {
	Data   []int  `json:"data"`
	Name   string `json:"name"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// JSONTileset ...
type JSONTileset struct {
	FirstGID   int    `json:"firstgid"`
	Name       string `json:"name"`
	Source     string `json:"image"`
	TileWidth  int    `json:"tilewidth"`
	TileHeight int    `json:"tileheight"`
	Width      int    `json:"imagewidth"`
	Height     int    `json:"imageheight"`
	Spacing    int    `json:"spacing"`
	Margin     int    `json:"margin"`
	Columns    int    `json:"columns"`
}

// --- JSON PARSER

// JSONMapParser ...
type JSONMapParser struct {
	tilesize int
	width    int
	height   int
}

// NewJSONMapParser ...
func NewJSONMapParser() *JSONMapParser {
	mp := &JSONMapParser{}
	return mp
}

// load data ...
func (mp *JSONMapParser) loadData(filename string) JSONMap {
	var err error

	file, err := os.Open(filename)
	defer file.Close()
	checkError(err)

	bv, err := ioutil.ReadAll(file)
	checkError(err)

	var data JSONMap
	err = json.Unmarshal(bv, &data)
	checkError(err)

	return data
}

// ParseLevel ...
func (mp *JSONMapParser) ParseLevel(filename string) *Level {
	data := mp.loadData(filename)

	level := NewLevel()

	// basic map settings
	mp.tilesize = data.Tilewidth
	mp.width = data.Width
	mp.height = data.Height

	// update levels []*Tileset slice
	for _, v := range data.Set {
		mp.parseTileSets(v, level)
	}

	// update levels []ILayer slice
	for _, v := range data.Layers {
		mp.parseTileLayers(v, level)
	}

	return level
}

// parse tile sets
// jTileset - slice of parsed data
// level * - pointer to level data
func (mp *JSONMapParser) parseTileSets(jTileset JSONTileset, level *Level) {
	// TODO - parse tile set json
	// parse tilesets from json file and create a new tile set from that data
	// then push that data to a slice for use
	// load in texture
	STextureManager.Load(jTileset.Source, jTileset.Name, STheGame.GetRenderer())

	// new tile set from data
	tileset := NewTileset(
		jTileset.FirstGID,   // groidid
		jTileset.Width,      // image width
		jTileset.Height,     // image height
		jTileset.TileWidth,  // tile width
		jTileset.TileHeight, // tile height
		jTileset.Spacing,    // spacing between tiles
		jTileset.Margin,     // margin between tiles
		jTileset.Columns,    // number of columns
		jTileset.Name)       // name of the tile set

	gologger.SLogger.Println("DATA FROM PARSE TILE LAYER::", tileset)

	// push to levels tile set slice
	level.AppendToTileSet(tileset)
}

// parse tile layers
// jlayers - slice of parsed data
// level * - pointer to level data
func (mp *JSONMapParser) parseTileLayers(jLayer JSONLayers, level *Level) {
	// TODO - parse tile layer json
	// in the book the xml data is compressed via zlib base64
	// so needs to uncompressed
	// but im using json were its not compressed so skipping the
	// uncompress part!
	// ---
	// also in the book a pointer to a vector pointer is used to add to
	// levels tile layers and also get tile set

	tilelayer := NewTileLayer(mp.tilesize, level.GetTileSet())

	// make empty tile data 2d slice
	data := make([][]int, mp.height)
	for d := 0; d < mp.width; d++ {
		data[d] = make([]int, mp.width)
	}

	// add data
	for rows := 0; rows < mp.height; rows++ {
		for cols := 0; cols < mp.width; cols++ {
			data[rows][cols] = jLayer.Data[rows*mp.width+cols]
		}
	}

	gologger.SLogger.Println("DATA FROM PARSE TILE LAYER::", data)

	tilelayer.SetTileIDs(data)

	// push to levels tile layer slice
	level.AppendToTileLayer(tilelayer)
}
