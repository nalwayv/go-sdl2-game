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
}

// XMLTextures ...
type XMLTextures struct {
	Filename string `xml:"filename"`
	ID       string `xml:"id"`
}

// XMLObjects ...
type XMLObjects struct {
	Width      int32  `xml:"width"`
	Height     int32  `xml:"height"`
	X          int32  `xml:"x"`
	Y          int32  `xml:"y"`
	ID         string `xml:"id"`
	NumFrames  int    `xml:"numframes"`
	CallBackID int    `xml:"callbackid"`
	AnimSpeed  int    `xml:"animspeed"`
	Type       string `xml:"type,attr"`
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

		sp.parseObjects(data.Menu.XMLObjects, o)
		sp.parseTextures(data.Menu.XMLTextures, t)
	}
}

// ParseTextures ...
func (sp *StateParser) parseTextures(textures []XMLTextures, t *[]string) {
	for _, v := range textures {
		STextureManager.Load(v.Filename, v.ID, STheGame.GetRenderer())

		*t = append(*t, v.ID)

		gologger.SLogger.Println("Pushed onto textureid slice", v.ID)
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
