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
}

// OnEnter ...
func (ps *PlayState) OnEnter() bool {
	fmt.Println("enter play state")

	sp := NewStateParser()
	sp.ParseState("data/tmp.xml", PlayID, &ps.objects, &ps.textureIDs)

	return true
}

// OnExit ...
func (ps *PlayState) OnExit() bool {
	fmt.Println("exit play state")

	for _, v := range ps.textureIDs {
		STextureManager.ClearFromTextureMap(v)
	}

	return true
}

// GetStateID ... get player id
func (ps *PlayState) GetStateID() string {
	return PlayID
}
