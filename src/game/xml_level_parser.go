package game

import (
	"encoding/xml"
	"io/ioutil"
	"os"
)

// --- XML Parser

// XMLFoo ...
type XMLMap struct{
	XMLName xml.Name `xml:"map"`
	Width int `xml:"width,attr"`
	Height int `xml:"height,attr"`
	TileSize int `xml:"tilewidth,attr"`
	XMLTileset struct {

	}`xml:"tileset"`

	XMLLayer struct{

		XMLData struct{

		}`xml`

	}`xml:"layer"`
}


// LevelParser ...
type LevelParser struct{}

// NewLevelParser ...
func NewLevelParser()*LevelParser{
	lp := &LevelParser{}
	return lp
}

func (lp *LevelParser)loadData(fileName string)XMLMap{
	var err error

	file,err := os.Open(fileName)
	defer file.Close()
	checkError(err)

	bv,err := ioutil.ReadAll(file)
	checkError(err)

	var data XMLMap
	err = xml.Unmarshal(bv, &data)
	checkError(err)

	return data
}

// parse level tile sets ...
func (lp *LevelParser)parseSets(){}

// parse level tile layers ...
func (lp *LevelParser)parseLayers(){}
