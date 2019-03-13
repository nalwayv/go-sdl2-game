package test

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"../../game"
)

func setUp() (*game.XMLStates, bool) {
	var err error
	var data game.XMLStates

	file, err := os.Open("../../../data/tmp.xml")
	if err != nil {
		fmt.Println("file error")
		return nil, false
	}
	defer file.Close()

	bv, _ := ioutil.ReadAll(file)
	xml.Unmarshal(bv, &data)

	return &data, true
}

func TestParserMenu(t *testing.T) {
	data, ok := setUp()

	if !ok {
		t.Error("error parsing file")
	}

	if len(data.Menu.XMLTextures) == 0 {
		t.Error("did not parse any textures")
	}

	if len(data.Menu.XMLTextures) < 2 {
		t.Error("did not parse all textures")
	}

	if len(data.Menu.XMLObjects) == 0 {
		t.Error("did not parse any objects")
	}

	if len(data.Menu.XMLObjects) < 2 {
		t.Error("did not parse all objects")
	}

}
