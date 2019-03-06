package game

// StateMachine ... game state machine
type StateMachine struct {
	gameState []IGameState // implement IGameState interface
}

// NewStateMachine ...
func NewStateMachine() *StateMachine {
	sm := &StateMachine{}
	sm.gameState = make([]IGameState, 0)
	return sm
}

// PushState ...
func (sm *StateMachine) PushState(state IGameState) {
	sm.gameState = append(sm.gameState, state)
	n := len(sm.gameState) - 1
	sm.gameState[n].OnEnter()
}

// PopState ...
func (sm *StateMachine) PopState() {
	// exit if empty
	if len(sm.gameState) == 0 {
		return
	}

	n := len(sm.gameState) - 1
	if sm.gameState[n].OnExit() {
		// remove last element
		sm.gameState = append(sm.gameState[n:], sm.gameState[n+1:]...)
	}
}

// ChangeState ...
func (sm *StateMachine) ChangeState(state IGameState) {
	// remove old
	if len(sm.gameState) != 0 {
		n := len(sm.gameState) - 1

		// same id
		if sm.gameState[n].GetStateID() == state.GetStateID() {
			// do nothing
			return
		}

		if sm.gameState[n].OnExit() {
			// pop last
			sm.gameState = append(sm.gameState[n:], sm.gameState[n+1:]...)
		}
	}

	// add new
	sm.gameState = append(sm.gameState, state)
	n := len(sm.gameState) - 1
	sm.gameState[n].OnEnter()
}

// Update ...
func (sm *StateMachine) Update() {
	if len(sm.gameState) != 0 {
		n := len(sm.gameState) - 1
		sm.gameState[n].Update()
	}
}

// Render ...
func (sm *StateMachine) Render() {
	if len(sm.gameState) != 0 {
		n := len(sm.gameState) - 1
		sm.gameState[n].Render()
	}
}
