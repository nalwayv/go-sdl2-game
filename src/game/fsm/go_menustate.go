package fsm

import "fmt"

// MenuState ...
type MenuState struct {
	state  *GameState // inherits
	menuID string
}

func newMenuState() *MenuState {
	ms := &MenuState{}
	ms.state = newGameState()
	ms.menuID = "MENU"
	return ms
}

// Update ...
func (ms MenuState) Update() {
	// ...
}

// Render ...
func (ms MenuState) Render() {
	// ...
}

// OnEnter ...
func (ms MenuState) OnEnter() bool {
	fmt.Println("enter menu state")
	return true
}

// OnExit ...
func (ms MenuState) OnExit() bool {
	fmt.Println("exit menu state")
	return true
}

// GetStateID ...
func (ms MenuState) GetStateID() string {
	return ms.menuID
}
