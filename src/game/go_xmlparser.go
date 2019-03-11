package game

import (
	"encoding/xml"
	"io/ioutil"
	"os"
)

// PARSER
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

// NewXMLStates ...
func NewXMLStates(fileName string) *XMLStates {
	states := load(fileName)
	return states
}

// Load ...
func load(fileName string) *XMLStates {

	file, err := os.Open(fileName)
	checkError(err)
	defer file.Close()

	bv, err := ioutil.ReadAll(file)
	checkError(err)

	var res XMLStates
	xml.Unmarshal(bv, &res)

	return &res
}

// ---

// StateParser ...
type StateParser struct {
	stateID   string
	fileName  string
	data      *XMLStates
	textureID []string
	objects   []IGameObject
}

// NewStateParser ...
func NewStateParser() *StateParser {
	sp := &StateParser{}
	sp.data = NewXMLStates("data/tmp.xml")
	sp.textureID = make([]string, 0)
	sp.objects = make([]IGameObject, 0)
	return sp
}

func (sp *StateParser) parseTextures() {
	for _, v := range sp.data.Menu.XMLTextures {
		STextureManager.Load(v.Filename, v.ID, STheGame.GetRenderer())
		sp.textureID = append(sp.textureID, v.ID)
	}
}

func (sp *StateParser) parseObjects() {
	for _, v := range sp.data.Menu.XMLObjects {
		// create object of type
		obj, err := STheGameObjFactory.Create(v.Type)
		checkError(err)

		// set its params
		obj.Load(NewParams(v.X, v.Y, v.Width, v.Height, v.ID, v.NumFrames, v.CallBackID, v.AnimSpeed))

		// store
		sp.objects = append(sp.objects, obj)
	}
}
