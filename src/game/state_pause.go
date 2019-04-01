package game

import "fmt"

/*
IMenuState
---
 - SetCallBacks([]Callback)
 - IGameState
   ---
   - Update()
   - Render()
   - OnEnter() bool
   - OnExit() bool
   - GetStateID() string
**/

// PauseID ... id for this object used for parsing state info
const PauseID string = "pause"

// PauseState ...
type PauseState struct {
	objects    []IGameObject
	textureIDs []string
	callbacks  MCallbacks
}

// NewPauseState ...
func NewPauseState() *PauseState {
	ps := &PauseState{}

	ps.objects = make([]IGameObject, 0)
	ps.textureIDs = make([]string, 0)
	ps.callbacks = make(MCallbacks, 0)

	return ps
}

// Update ...
func (ps *PauseState) Update() {
	for _, v := range ps.objects {
		v.Update()
	}
}

// Render ...
func (ps *PauseState) Render() {
	for _, v := range ps.objects {
		v.Draw()
	}
}

// OnEnter ...
func (ps *PauseState) OnEnter() bool {
	fmt.Println("enter pause state")

	sp := NewJSONStateParser()
	sp.ParseState("data/data.json", PauseID, &ps.objects, &ps.textureIDs)

	// button callback functions
	// starts from 1 so 0 is nil
	ps.callbacks = append(ps.callbacks, nil)
	ps.callbacks = append(ps.callbacks, pauseToMain)
	ps.callbacks = append(ps.callbacks, resume)
	ps.SetCallBacks(ps.callbacks)

	return true
}

// OnExit ...
func (ps *PauseState) OnExit() bool {
	fmt.Println("exit pause state")

	for _, v := range ps.textureIDs {
		STextureManager.ClearFromTextureMap(v)
	}

	SInputHandler.Reset()

	return true
}

// SetCallBacks ...
func (ps *PauseState) SetCallBacks(cb []Callback) {
	// are of type IGameObject
	for _, v := range ps.objects {
		switch v.(type) {
		// if type menubutton
		case *MenuButton:
			button := v.(*MenuButton)

			// set callback based on button loaded id
			// - 1:: pauseToMenu
			// - 2:: resume
			cb := ps.callbacks[button.GetCallBackID()]
			button.SetCallBack(cb)
		}
	}
}

// GetStateID ...
func (ps PauseState) GetStateID() string {
	return PauseID
}

// --- CallBacks

// go back to main menu
func pauseToMain() {
	STheGame.GetStateMachine().ChangeState(NewMenuState())
}

// resume previous state
func resume() {
	STheGame.GetStateMachine().PopState()
}
