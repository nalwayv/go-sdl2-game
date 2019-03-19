package game

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"../gologger"
)

// --- JSON DATA

// JSONStates ...
type JSONStates struct {
	State struct {
		Menu struct {
			JSONTextures []JSONTextures `json:"textures"`
			JSONObjects  []JSONObjects  `json:"objects"`
		} `json:"menu"`

		Play struct {
			JSONTextures []JSONTextures `json:"textures"`
			JSONObjects  []JSONObjects  `json:"objects"`
		} `json:"play"`

		Pause struct {
			JSONTextures []JSONTextures `json:"textures"`
			JSONObjects  []JSONObjects  `json:"objects"`
		} `json:"pause"`

		GameOver struct {
			JSONTextures []JSONTextures `json:"textures"`
			JSONObjects  []JSONObjects  `json:"objects"`
		} `json:"gameover"`
	} `json:"states"`
}

// JSONTextures ...
type JSONTextures struct {
	FileName string `json:"filename"`
	ID       string `json:"id"`
}

// JSONObjects ...
type JSONObjects struct {
	Type       string `json:"type"`
	X          int32  `json:"x"`
	Y          int32  `json:"y"`
	Width      int32  `json:"width"`
	Height     int32  `json:"height"`
	ID         string `json:"id"`
	NumFrames  int    `json:"numframes"`
	CallBackID int    `json:"callbackid"`
	AnimSpeed  int    `json:"animspeed"`
}

// --- JSON PARSER

// JSONStateParser ...
type JSONStateParser struct{}

// NewJSONStateParser ...
func NewJSONStateParser() *JSONStateParser {
	gologger.SLogger.Println("init json state parser")

	jsp := &JSONStateParser{}

	return jsp
}

// loadData ...
func (jsp *JSONStateParser) loadData(filename string) JSONStates {
	var err error

	file, err := os.Open(filename)
	defer file.Close()
	checkError(err)

	bv, err := ioutil.ReadAll(file)
	checkError(err)

	var data JSONStates
	err = json.Unmarshal(bv, &data)
	checkError(err)

	return data
}

// ParseState ...
// filename	:: file with data
// stateID	:: id for parsing
// o		:: pointer to slice where data will be stored
// t		:: pointer to slice were data will be stored
func (jsp *JSONStateParser) ParseState(filename, stateID string, o *[]IGameObject, t *[]string) {
	data := jsp.loadData(filename)

	switch stateID {
	case "menu":
		gologger.SLogger.Println("Parsing Menu State")
		jsp.parseTextures(data.State.Menu.JSONTextures, t)
		jsp.parseObjects(data.State.Menu.JSONObjects, o)

	case "play":
		gologger.SLogger.Println("Parsing Play State")
		jsp.parseTextures(data.State.Play.JSONTextures, t)
		jsp.parseObjects(data.State.Play.JSONObjects, o)

	case "pause":
		gologger.SLogger.Println("Parsing Pause State")
		jsp.parseTextures(data.State.Pause.JSONTextures, t)
		jsp.parseObjects(data.State.Pause.JSONObjects, o)

	case "gameover":
		gologger.SLogger.Println("Parsing GameOver State")
		jsp.parseTextures(data.State.GameOver.JSONTextures, t)
		jsp.parseObjects(data.State.GameOver.JSONObjects, o)

	default:
		gologger.SLogger.Println("StateID not found")
	}
}

// parse textures
func (jsp *JSONStateParser) parseTextures(textures []JSONTextures, t *[]string) {
	for _, v := range textures {
		STextureManager.Load(v.FileName, v.ID, STheGame.GetRenderer())

		*t = append(*t, v.ID)

		gologger.SLogger.Println("Pushed onto textureIDs", v.ID)
	}
}

// parse objects
func (jsp *JSONStateParser) parseObjects(objects []JSONObjects, o *[]IGameObject) {
	for _, v := range objects {
		obj, err := STheGameObjFactory.Create(v.Type)

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

		*o = append(*o, obj)

		gologger.SLogger.Println("Created", v.Type)
	}
}
