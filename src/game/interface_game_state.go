package game


/*
* Info
* ---
* interface for game state classes
**/

// IGameState ... game state interface
type IGameState interface {
	Update()
	Render()
	OnEnter() bool
	OnExit() bool
	GetStateID() string
}
