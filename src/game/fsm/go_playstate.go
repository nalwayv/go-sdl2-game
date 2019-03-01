package fsm

import "fmt"

// PlayState ...
type PlayState struct {
	state  *GameState
	playID string
}

func newPlayState() *PlayState {
	ps := &PlayState{}
	ps.state = newGameState()
	ps.playID = "PLAY STATE"
	return ps
}

// Update ...
func (ps PlayState) Update() {
	// ...
}

// Render ...
func (ps PlayState) Render() {
	// ...
}

// OnEnter ...
func (ps PlayState) OnEnter() bool {
	fmt.Println("enter play state")
	return true
}

// OnExit ...
func (ps PlayState) OnExit() bool {
	fmt.Println("exit play state")
	return true
}

// GetStateID ...
func (ps PlayState) GetStateID() string {
	return ps.playID
}
