package game

// TODO

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type JsonObj struct {
	State `json:"States"`

	TextureIDs []string
}

type State struct {
	Menu `json:"Menu"`
}

type Menu struct {
	Textures []Textures `json:"Textures"`
	Objects  []Objects  `json:"Objects"`
}

type Textures struct {
	Filename string `json:"filename"`
	ID       string `json:"ID"`
}

type Objects struct {
	Objtype   string `json:"objtype"`
	X         int32  `json:"x"`
	Y         int32  `json:"y"`
	Width     int32  `json:"width"`
	Height    int32  `json:"height"`
	TextureID string `json:"textureID"`
}

// NewJSONObj ...
func NewJSONObj() *JsonObj {
	jsonfile := load()
	jsonfile.TextureIDs = make([]string, 0)
	return jsonfile
}

func load() *JsonObj {
	var err error

	jsonfile, err := os.Open("data/data.json")
	checkError(err)
	defer jsonfile.Close()

	bv, err := ioutil.ReadAll(jsonfile)
	checkError(err)

	var result JsonObj
	json.Unmarshal(bv, &result)

	return &result

}

// ParseTextures ... parse textures from loaded json file
func (j *JsonObj) ParseTextures() {
	for _, v := range j.Textures {
		j.TextureIDs = append(j.TextureIDs, v.ID)

		STextureManager.Load(v.Filename, v.ID, STheGame.GetRenderer())
	}
}

// ParseObjects ...
func (j *JsonObj) ParseObjects() {
	// for _, v := range j.Objects {
	//
	// }
}
