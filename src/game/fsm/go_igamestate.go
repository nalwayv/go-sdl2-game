package fsm

// IGameState ... game state interface
type IGameState interface {
	Update()
	Render()
	OnEnter() bool
	OnExit() bool
	GetStateID() string
}
