package game

/*
Implements IGameState interface.

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

// PlayState ...
type PlayState struct {
	playID      string
	gameObjects []IGameObject
}

// NewPlayState ...
func NewPlayState() *PlayState {
	ps := &PlayState{}
	ps.playID = "PLAY"
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

	for _, v := range ps.gameObjects {
		v.Update()
	}
}

// Render ...
func (ps *PlayState) Render() {
	for _, v := range ps.gameObjects {
		v.Draw()
	}
}

// OnEnter ...
func (ps *PlayState) OnEnter() bool {
	fmt.Println("enter play state")

	// load textures
	STextureManager.Load("assets/helicopter.png", "player", STheGame.GetRenderer())

	// new player
	player := NewPlayer()

	// add to game objects slice
	ps.gameObjects = append(ps.gameObjects, player)

	return true
}

// OnExit ...
func (ps *PlayState) OnExit() bool {
	fmt.Println("exit play state")

	for _, v := range ps.gameObjects {
		v.Clean()
	}

	ps.gameObjects = nil

	STextureManager.ClearFromTextureMap("player")

	return true
}

// GetStateID ... get player id
func (ps PlayState) GetStateID() string {
	return ps.playID
}
