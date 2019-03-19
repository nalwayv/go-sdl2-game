package game

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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
	Name   string `json:"name"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Data   []int  `json:"data"`
}

// JSONTileset ...
type JSONTileset struct {
	FirstGID   int    `json:"firstgid"`
	Name       string `json:"name"`
	Source     string `json:"source"`
	TileWidth  int    `json:"tilewidth"`
	TileHeight int    `json:"tileheight"`
	Width      int    `json:"width"`
	Height     int    `json:"height"`
	Spacing    int    `json:"spacing"`
	Margin     int    `json:"margin"`
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
func (mp *JSONMapParser) ParseLevel(filename string) {
	// data:=loadData(filename)

}

// parse tile sets
func (mp *JSONMapParser) parseTileSets(ts []JSONTileset) {
	for _, v := range ts {
		// new tile set from data
		t := NewTileset()
		t.Width = v.Width
		t.Height = v.Height
		t.GridID = v.FirstGID
		t.TileWidth = v.TileWidth
		t.TileHeight = v.TileHeight
		t.Spacing = v.Spacing
		t.Margin = v.Margin
		t.Name = v.Name
		t.NumColumns = t.Width / (t.TileWidth + t.Spacing)
		fmt.Println(t)
		// TODO
		// append t to tile set slice
	}
}

// parse tile layers
func (mp *JSONMapParser) parseTileLayers(tl []JSONLayers) {
	for _, v := range tl {
		// TODO
		// new tile layer
		//p := NewTileLayer(mp.tilesize, nil)
		// tile data
		// empty matrix
		data := make([][]int, mp.height)
		for d := 0; d < mp.width; d++ {
			data[d] = make([]int, mp.width)
		}
		// add data
		for row := 0; row < mp.height; row++ {
			for col := 0; col < mp.width; col++ {
				data[row][col] = v.Data[row*mp.width+col]
			}
		}
		fmt.Println(data)
		// TODO
		// append data to slice
	}
}
