package fsm

// GameState ...
type GameState struct {
	id string
}

func newGameState() *GameState {
	gs := &GameState{}
	gs.id = "GAME STATE"
	return gs
}

// Update ...
func (gs GameState) Update() {
}

// Render ...
func (gs GameState) Render() {
}

// OnEnter ...
func (gs GameState) OnEnter() bool {
	return true
}

// OnExit ...
func (gs GameState) OnExit() bool {
	return true
}

// GetStateID ...
func (gs GameState) GetStateID() string {
	return gs.id
}
