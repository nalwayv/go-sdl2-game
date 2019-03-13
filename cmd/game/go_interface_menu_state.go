package game

/*
IGameState interface
---

-Update()
-Render()
-OnEnter() bool
-OnExit() bool
-GetStateID() string
*/

// Callback ...
type Callback func()

// MCallbacks ...
type MCallbacks []Callback

// IMenuState ...
type IMenuState interface {
	IGameState
	SetCallBacks([]Callback)
}
