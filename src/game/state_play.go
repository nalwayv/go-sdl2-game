package game

/*
*IGameState
---

- Update()
- Render()
- OnEnter() bool
- OnExit() bool
- GetStateID() string
*/

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

// PlayID ...  id for this object used for parsing state info
const PlayID string = "play"

// PlayState ...
type PlayState struct {
	objects    []IGameObject
	textureIDs []string
	pLevel     *Level
}

// NewPlayState ...
func NewPlayState() *PlayState {
	ps := &PlayState{}
	ps.objects = make([]IGameObject, 0)
	ps.textureIDs = make([]string, 0)
	return ps
}

// Update ...
func (ps *PlayState) Update() {

	// push temp pause state onto fsm stack
	if SInputHandler.IsKeyDown(sdl.SCANCODE_ESCAPE) {
		STheGame.GetStateMachine().PushState(NewPauseState())
	}

	if SInputHandler.IsKeyDown(sdl.SCANCODE_Q) {
		STheGame.GetStateMachine().PushState(NewGameOverState())
	}

	for _, v := range ps.objects {
		v.Update()
	}
}

// Render ...
func (ps *PlayState) Render() {
	for _, v := range ps.objects {
		v.Draw()
	}

	ps.pLevel.Render()
}

// OnEnter ...
func (ps *PlayState) OnEnter() bool {
	fmt.Println("enter play state")

	// obj
	sp := NewJSONStateParser()
	sp.ParseState("data/data.json", PlayID, &ps.objects, &ps.textureIDs)

	// level
	lp := NewJSONMapParser()
	ps.pLevel = lp.ParseLevel("data/map.json")

	return true
}

// OnExit ...
func (ps *PlayState) OnExit() bool {
	fmt.Println("exit play state")

	// for _, v := range ps.textureIDs {
	//     STextureManager.ClearFromTextureMap(v)
	// }

	for _, v := range ps.textureIDs {
		err := STextureManager.ClearFromTextureMap(v)
		checkError(err)
	}

	return true
}

// GetStateID ... get player id
func (ps *PlayState) GetStateID() string {
	return PlayID
}
