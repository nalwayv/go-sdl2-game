package game

import (
	"encoding/xml"
	"fmt"
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
type StateParser struct {
	textureID []string
	objects   []IGameObject
}

// NewStateParser ...
func NewStateParser() *StateParser {
	sp := &StateParser{}

	sp.textureID = make([]string, 0)
	sp.objects = make([]IGameObject, 0)

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
func (sp *StateParser) ParseState(fileName, stateID string) {
	data := sp.loadData(fileName)

	if stateID == "menu" {
		sp.parseTextures(data.Menu.XMLTextures)
		sp.parseObjects(data.Menu.XMLObjects)
	}

}

// ParseTextures ...
func (sp *StateParser) parseTextures(textures []XMLTextures) {
	for _, v := range textures {
		STextureManager.Load(v.Filename, v.ID, STheGame.GetRenderer())
		sp.textureID = append(sp.textureID, v.ID)
	}

}

// ParseObjects ...
func (sp *StateParser) parseObjects(objects []XMLObjects) {
	for _, v := range objects {
		// create object of type
		gologger.SLogger.Println("Creating", v.Type)

		// create obj of type and set its params
		obj, err := STheGameObjFactory.Create(v.Type)
		gologger.SLogger.Println("Created", v.Type)

		checkError(err)
		obj.Load(NewParams(
			v.X,
			v.Y,
			v.Width,
			v.Height,
			v.ID,
			v.NumFrames,
			v.CallBackID,
			v.AnimSpeed))

		gologger.SLogger.Println("Added params for", v.Type)

		sp.objects = append(sp.objects, obj)
		fmt.Println("added")

	}
}

// GetParsedObjects ...
func (sp *StateParser) GetParsedObjects() []IGameObject {
	return sp.objects
}

// GetParsedTextureIDs ...
func (sp *StateParser) GetParsedTextureIDs() []string {
	return sp.textureID
}
