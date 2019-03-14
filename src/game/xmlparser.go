package game

import (
	"encoding/xml"
	"io/ioutil"
	"os"

	"../gologger"
)

// --- XML Data

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

	Pause struct {
		XMLTextures []XMLTextures `xml:"textures>texture"`
		XMLObjects  []XMLObjects  `xml:"objects>object"`
	} `xml:"pause"`

	GameOver struct {
		XMLTextures []XMLTextures `xml:"textures>texture"`
		XMLObjects  []XMLObjects  `xml:"objects>object"`
	} `xml:"gameover"`
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

// --- XML Parser

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
// filename - string
// stateID - string - id for the state so its textures/objects are parsed
// o - *[] IGameObject - 'game objects' :: pointing back to slice with data that will be appended to it from parser
// t - *[] string - 'texture ids' :: pointing back to slice with data that will be appended to it from parser
func (sp *StateParser) ParseState(fileName, stateID string, o *[]IGameObject, t *[]string) {
	data := sp.loadData(fileName)

	switch stateID {
	case "menu":
		gologger.SLogger.Println("Parsing Menu State")
		sp.parseTextures(data.Menu.XMLTextures, t)
		sp.parseObjects(data.Menu.XMLObjects, o)

	case "play":
		gologger.SLogger.Println("Parsing Play State")
		sp.parseTextures(data.Play.XMLTextures, t)
		sp.parseObjects(data.Play.XMLObjects, o)

	case "pause":
		gologger.SLogger.Println("Parsing Pause State")
		sp.parseTextures(data.Pause.XMLTextures, t)
		sp.parseObjects(data.Pause.XMLObjects, o)

	case "gameover":
		gologger.SLogger.Println("Parsing GameOver State")
		sp.parseTextures(data.GameOver.XMLTextures, t)
		sp.parseObjects(data.GameOver.XMLObjects, o)

	default:
		gologger.SLogger.Println("StateID not found")
	}
}

// ParseTextures ...
func (sp *StateParser) parseTextures(textures []XMLTextures, t *[]string) {
	for _, v := range textures {
		STextureManager.Load(v.Filename, v.ID, STheGame.GetRenderer())

		// append to states texture id list
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

		// append to states object list
		*o = append(*o, obj)

		gologger.SLogger.Println("Created", v.Type)
	}
}
