package game

/*
* Info
* ___
* An interface that implements another interface as the buttons use callback functions to change states
*
* IGameState interface
* ---
*
* -Update()
* -Render()
* -OnEnter() bool
* -OnExit() bool
* -GetStateID() string
**/

// Callback ...
type Callback func()

// MCallbacks ...
type MCallbacks []Callback

// IMenuState ...
type IMenuState interface {
	IGameState
	SetCallBacks([]Callback)
}
