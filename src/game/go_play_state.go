package game

import "fmt"

/*
	*IGameState*
	- Update()
	- Render()
	- OnEnter() bool
	- OnExit() bool
	- GetStateID() string
*/

// PlayState ...
type PlayState struct {
	playID      string
	gameObjects []IGameObject
}

// NewPlayState ...
func NewPlayState() *PlayState {
	ps := &PlayState{}
	ps.playID = "PLAY"
	return ps
}

// Update ...
func (ps *PlayState) Update() {
	for _, v := range ps.gameObjects {
		v.Update()
	}
}

// Render ...
func (ps *PlayState) Render() {
	for _, v := range ps.gameObjects {
		v.Draw()
	}
}

// OnEnter ...
func (ps *PlayState) OnEnter() bool {
	fmt.Println("enter play state")

	STextureManager.Load("assets/handheld.png", "player", STheGame.GetRenderer())

	player := NewPlayer(NewParams(0, 0, 96, 96, "animate"))

	ps.gameObjects = append(ps.gameObjects, player)

	return true
}

// OnExit ...
func (ps *PlayState) OnExit() bool {
	fmt.Println("exit play state")

	for _, v := range ps.gameObjects {
		v.Clean()
	}

	ps.gameObjects = nil

	STextureManager.ClearFromTextureMap("player")

	return true
}

// GetStateID ...
func (ps PlayState) GetStateID() string {
	return ps.playID
}
