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
	pLevel *Level
}

// NewPlayState ...
func NewPlayState() *PlayState {
	ps := &PlayState{}
	return ps
}

// Update ...
func (ps *PlayState) Update() {

	// push temp pause state onto fsm stack
	if SInputHandler.IsKeyDown(sdl.SCANCODE_ESCAPE) {
		STheGame.GetStateMachine().PushState(NewPauseState())
	}

	ps.pLevel.Update()
}

// Render ...
func (ps *PlayState) Render() {
	ps.pLevel.Render()
}

// OnEnter ...
func (ps *PlayState) OnEnter() bool {
	fmt.Println("enter play state")

	// level's info
	lp := NewJSONMapParser()
	ps.pLevel = lp.ParseLevel("data/map.json")

	return true
}

// OnExit ...
func (ps *PlayState) OnExit() bool {
	fmt.Println("exit play state")

	return true
}

// GetStateID ... get player id
func (ps *PlayState) GetStateID() string {
	return PlayID
}
