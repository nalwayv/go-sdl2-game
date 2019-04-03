package game

/*
Info
---
Data parsed from the json file generated from 'Tiled'
then passed onto the current level Level
to be stored within ever its 'tilesets' or 'layers' slice
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
	Tileheight int `json:"tileheight"`
	Tilewidth  int `json:"tilewidth"`
	Width      int `json:"width"`
	Height     int `json:"height"`

	TileSets   []JSONTileset   `json:"tilesets"`
	Layers     []JSONLayers    `json:"layers"`
	Properties []JSONProp      `json:"properties"`
	ObjGroups  []JSONObjGroups `json:"objectgroups"`
}

// JSONLayers ...
type JSONLayers struct {
	Data   []int  `json:"data"`
	Name   string `json:"name"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
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
	Columns    int    `json:"columns"`
}

// JSONProp ...
type JSONProp struct {
	Name   string `json:"name"`
	Source string `json:"source"`
}

// JSONObjGroups ...
type JSONObjGroups struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	X       int       `json:"x"`
	Y       int       `json:"y"`
	Objects []JSONObj `json:"objects"`
}

// JSONObj ...
type JSONObj struct {
	Name           string `json:"name"`
	Type           string `json:"type"`
	TexID          string `json:"id"`
	X              int32  `json:"x"`
	Y              int32  `json:"y"`
	Width          int32  `json:"width"`
	Height         int32  `json:"height"`
	NumberOfFrames int    `json:"numframes"`
	AnimSpeed      int    `json:"animspeed"`
	CallBackID     int    `json:"callbackid"`
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
	for _, v := range data.TileSets {
		mp.parseTileSets(v, level)
	}

	// update levels []ILayer slice
	for _, v := range data.Layers {
		mp.parseTileLayers(v, level)
	}

	for _, v := range data.Properties {
		mp.parseTextures(v)
	}

	for _, v := range data.ObjGroups {
		mp.parseObjLayers(v, level)
	}

	return level
}

// parse tile sets ...
func (mp *JSONMapParser) parseTileSets(jTileset JSONTileset, level *Level) {
	// load map texture from parse data
	STextureManager.Load(jTileset.Source, jTileset.Name, STheGame.GetRenderer())

	// new tile set from data
	tileset := NewTileset(
		jTileset.FirstGID,   // gridID
		jTileset.Width,      // image width
		jTileset.Height,     // image height
		jTileset.TileWidth,  // tile width
		jTileset.TileHeight, // tile height
		jTileset.Spacing,    // spacing between tiles
		jTileset.Margin,     // margin between tiles
		jTileset.Columns,    // number of columns
		jTileset.Name)       // name of the tile set

	gologger.SLogger.Println("Created new Tileset")

	// push to levels tile set slice
	level.AppendToTileSet(tileset)
}

// parse tile layers ...
func (mp *JSONMapParser) parseTileLayers(jLayer JSONLayers, level *Level) {
	// new texture layer to store info
	tilelayer := NewTileLayer(mp.tilesize, level.GetTileSet())

	// empty 2d slice to hold tile map data
	data := make([][]int, 0)
	for i := 0; i < mp.height; i++ {
		data = append(data, make([]int, mp.width))
	}

	// add data
	for rows := 0; rows < mp.height; rows++ {
		for cols := 0; cols < mp.width; cols++ {

			// convert to 2d map coords
			tileid := (rows * mp.width) + cols

			data[rows][cols] = jLayer.Data[tileid]
		}
	}

	gologger.SLogger.Println("Created new Tilelayer")

	tilelayer.SetTileIDs(data)

	// push to levels tile layer slice
	level.AppendToLayer(tilelayer)
}

// parse tile texture ...
func (mp *JSONMapParser) parseTextures(jProp JSONProp) {
	// get textures for objects in current level
	STextureManager.Load(jProp.Source, jProp.Name, STheGame.GetRenderer())

	gologger.SLogger.Println("Pushed onto textureIDs", jProp.Source)
}

// parse object layers
func (mp *JSONMapParser) parseObjLayers(jGroups JSONObjGroups, level *Level) {
	ol := NewObjectLayer()

	for _, v := range jGroups.Objects {
		// create obj based on type
		obj, err := STheGameObjFactory.Create(v.Type)

		checkError(err)

		obj.Load(NewParams(v.X, v.Y, v.Width, v.Height, v.TexID, v.NumberOfFrames, v.CallBackID, v.AnimSpeed))

		gologger.SLogger.Println("Created obj from pared data of type", v.Type)

		ol.PushOntoGameObj(obj)

		level.AppendToLayer(ol)
	}
}
