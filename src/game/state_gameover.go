package game

/*
*IGameState
---
	- Update()
	- Render()
	- OnEnter() bool
	- OnExit() bool
	- GetStateID() string

*IMenuState
---

	- SetCallBacks([]Callback)
*/

import "fmt"

// GameOverID ...
const GameOverID string = "gameover"

// GameOverState ...
type GameOverState struct {
	objects    []IGameObject
	textureIDs []string
	callbacks  MCallbacks
}

// NewGameOverState ...
func NewGameOverState() *GameOverState {
	gs := &GameOverState{}

	gs.objects = make([]IGameObject, 0)
	gs.textureIDs = make([]string, 0)
	gs.callbacks = make(MCallbacks, 0)

	return gs
}

// Update ...
func (gs *GameOverState) Update() {
	for _, v := range gs.objects {
		v.Update()
	}
}

// Render ...
func (gs *GameOverState) Render() {
	for _, v := range gs.objects {
		v.Draw()
	}
}

// OnEnter ...
func (gs *GameOverState) OnEnter() bool {
	fmt.Println("enter gameover state")

	sp := NewStateParser()
	sp.ParseState("data/tmp.xml", GameOverID, &gs.objects, &gs.textureIDs)

	gs.callbacks = append(gs.callbacks, nil)
	gs.callbacks = append(gs.callbacks, gameoverToMenu)
	gs.callbacks = append(gs.callbacks, restart)
	gs.SetCallBacks(gs.callbacks)

	return true
}

// OnExit ...
func (gs *GameOverState) OnExit() bool {

	fmt.Println("exit gameover state")

	for _, v := range gs.textureIDs {
		STextureManager.ClearFromTextureMap(v)
	}

	return true
}

// GetStateID ... get state id
func (gs *GameOverState) GetStateID() string {
	return GameOverID
}

// SetCallBacks ...
func (gs *GameOverState) SetCallBacks(cb []Callback) {
	// are of type IGameObject
	for _, v := range gs.objects {
		switch v.(type) {
		// if type menubutton
		case *MenuButton:
			button := v.(*MenuButton)

			// set callback based on button loaded id
			// - 1:: gameoverToMenu
			// - 2:: restart
			cb := gs.callbacks[button.GetCallBackID()]
			button.SetCallBack(cb)
		}
	}
}

// --- Callbacks

// go to main menu
func gameoverToMenu() {
	STheGame.GetStateMachine().ChangeState(NewMenuState())
}

// go to new play state
func restart() {
	STheGame.GetStateMachine().ChangeState(NewPlayState())
}
