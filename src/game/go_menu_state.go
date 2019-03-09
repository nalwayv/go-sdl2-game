package game

/*
Implements IGameState interface.

- Update()
- Render()
- OnEnter() bool
- OnExit() bool
- GetStateID() string
*/

import (
	"fmt"
)

// MenuState ...
type MenuState struct {
	menuID      string
	gameObjects []IGameObject
}

// NewMenuState ...
func NewMenuState() *MenuState {
	ms := &MenuState{}
	ms.menuID = "MENU"
	ms.gameObjects = make([]IGameObject, 0)
	return ms
}

// Update ...
func (ms MenuState) Update() {
	for i := range ms.gameObjects {
		ms.gameObjects[i].Update()
	}
}

// Render ...
func (ms MenuState) Render() {
	for i := range ms.gameObjects {
		ms.gameObjects[i].Draw()
	}
}

// OnEnter ...
func (ms *MenuState) OnEnter() bool {
	fmt.Println("enter menu state")

	// load textures
	STextureManager.Load("assets/button.png", "playbutton", STheGame.GetRenderer())
	STextureManager.Load("assets/exit.png", "exitbutton", STheGame.GetRenderer())

	// set buttons / functions
	playbutton := NewMenuButton(NewParams(100, 100, 400, 100, "playbutton", 0), func() {
		fmt.Println("PLAY BUTTON CLICKED")
		STheGame.GetStateMachine().ChangeState(NewPlayState())
	})

	exitbutton := NewMenuButton(NewParams(100, 300, 400, 100, "exitbutton", 0), func() {
		fmt.Println("EXIT BUTTON CLICKED")
		STheGame.Quit()
	})

	// add to gameobjects slice
	ms.gameObjects = append(ms.gameObjects, playbutton)
	ms.gameObjects = append(ms.gameObjects, exitbutton)

	return true
}

// OnExit ...
func (ms MenuState) OnExit() bool {
	var err error

	fmt.Println("exit menu state")

	for _, v := range ms.gameObjects {
		v.Clean()
	}

	ms.gameObjects = nil // clear slice

	err = STextureManager.ClearFromTextureMap("playbutton")
	checkError(err)

	err = STextureManager.ClearFromTextureMap("exitbutton")
	checkError(err)

	return true
}

// GetStateID ...
func (ms MenuState) GetStateID() string {
	return ms.menuID
}
