package game

import (
	"encoding/xml"
	"io/ioutil"
	"os"

	"../gologger"
)

// XML PARSE DATA ---

// XMLStates ...
type XMLStates struct {
	XMLName xml.Name `xml:"states"`
	Menu    struct {
		XMLTextures []XMLTextures `xml:"textures>texture"`
		XMLObjects  []XMLObjects  `xml:"objects>object"`
	} `xml:"menu"`

	Play struct {
		XMLTextures []XMLTextures `xml:"textures>texture"`
		XMLObjects  []XMLObjects  `xml:"objects>object"`
	} `xml:"play"`
}

// XMLTextures ...
type XMLTextures struct {
	Filename string `xml:"filename,attr"`
	ID       string `xml:"id,attr"`
}

// XMLObjects ...
type XMLObjects struct {
	Type       string `xml:"type,attr"`
	X          int32  `xml:"x,attr"`
	Y          int32  `xml:"y,attr"`
	Width      int32  `xml:"width,attr"`
	Height     int32  `xml:"height,attr"`
	ID         string `xml:"id,attr"`
	NumFrames  int    `xml:"numframes,attr"`
	CallBackID int    `xml:"callbackid,attr"`
	AnimSpeed  int    `xml:"animspeed,attr"`
}

// XML PARSE DATA ---

// StateParser ...
type StateParser struct{}

// NewStateParser ...
func NewStateParser() *StateParser {
	gologger.SLogger.Println("Init State Parser")

	sp := &StateParser{}
	return sp
}

// golang .xml file parser
// creates an object that stores data
func (sp *StateParser) loadData(fileName string) XMLStates {
	var err error

	file, err := os.Open(fileName)
	defer file.Close()
	checkError(err)

	bv, err := ioutil.ReadAll(file)
	checkError(err)

	var data XMLStates
	err = xml.Unmarshal(bv, &data)
	checkError(err)

	return data
}

// ParseState ...
// o - *[] IGameObject :: 'game objects' :: pointing back to slice with data that will be appended to it from parser
// t - *[] string :: 'texture ids' :: pointing back to slice with data that will be appended to it from parser
func (sp *StateParser) ParseState(fileName, stateID string, o *[]IGameObject, t *[]string) {
	data := sp.loadData(fileName)

	if stateID == "menu" {
		gologger.SLogger.Println("Parsing Menu State")
		sp.parseTextures(data.Menu.XMLTextures, t)
		sp.parseObjects(data.Menu.XMLObjects, o)
	}

	if stateID == "play" {
		gologger.SLogger.Println("Parsing Play State")
		sp.parseTextures(data.Play.XMLTextures, t)
		sp.parseObjects(data.Play.XMLObjects, o)
	}
}

// ParseTextures ...
func (sp *StateParser) parseTextures(textures []XMLTextures, t *[]string) {
	for _, v := range textures {
		STextureManager.Load(v.Filename, v.ID, STheGame.GetRenderer())

		*t = append(*t, v.ID)

		gologger.SLogger.Println("Pushed onto textureID slice", v.ID)
	}
}

// ParseObjects ...
func (sp *StateParser) parseObjects(objects []XMLObjects, o *[]IGameObject) {
	for _, v := range objects {
		// create obj of type and set its params
		obj, err := STheGameObjFactory.Create(v.Type)

		checkError(err)

		obj.Load(NewParams(v.X, v.Y, v.Width, v.Height, v.ID, v.NumFrames, v.CallBackID, v.AnimSpeed))

		*o = append(*o, obj)

		gologger.SLogger.Println("Created", v.Type)
	}
}
