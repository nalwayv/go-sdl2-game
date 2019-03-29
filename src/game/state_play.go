package game

/*
into
---
main game state

implements IGameState
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

// PlayID ...  id for parsing state info
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

	// TMP: switch to pause state
	if SInputHandler.IsKeyDown(sdl.SCANCODE_ESCAPE) {
		STheGame.GetStateMachine().PushState(NewPauseState())
	}

	// TMP switch to GameOver Menu
	// if SInputHandler.IsKeyDown(sdl.SCANCODE_Q) {
	//     STheGame.GetStateMachine().PushState(NewGameOverState())
	// }

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
