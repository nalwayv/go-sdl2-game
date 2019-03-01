package fsm

// StateMachine ... game state machine
type StateMachine struct {
	gameState []IGameState // implement IGameState interface
}

func newStateMachine() *StateMachine {
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
	// exit if empty
	if len(sm.gameState) == 0 {
		return
	}

	n := len(sm.gameState) - 1

	if sm.gameState[n].GetStateID() == state.GetStateID() {
		// do nothing
		return
	}

	// pop last append new
	if sm.gameState[n].OnExit() {
		sm.gameState = append(sm.gameState[n:], sm.gameState[n+1:]...)
	}

	sm.gameState = append(sm.gameState, state)
	sm.gameState[n].OnEnter()
}
