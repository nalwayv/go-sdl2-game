package game

/*
Implements IMenuState interface.
---

*IGameState
	- Update()
	- Render()
	- OnEnter() bool
	- OnExit() bool
	- GetStateID() string

*IMenuState
	- SetCallBacks([]Callback)
*/

import (
	"fmt"
)

// MenuID ... string if for object
const MenuID string = "menu"

// MainMenuState ...
type MainMenuState struct {
	//menuID     string
	objects    []IGameObject
	textureIDs []string
	callbacks  MCallbacks
}

// NewMenuState ...
func NewMenuState() *MainMenuState {
	ms := &MainMenuState{}
	//ms.menuID = "menu"
	ms.objects = make([]IGameObject, 0)
	ms.textureIDs = make([]string, 0)
	ms.callbacks = make(MCallbacks, 0)
	return ms
}

// Update ...
func (ms MainMenuState) Update() {
	for i := range ms.objects {
		ms.objects[i].Update()
	}
}

// Render ...
func (ms MainMenuState) Render() {
	for i := range ms.objects {
		ms.objects[i].Draw()
	}
}

// OnEnter ...
func (ms *MainMenuState) OnEnter() bool {
	fmt.Println("enter main menu state")

	sp := NewStateParser()
	sp.ParseState("data/tmp.xml", MenuID, &ms.objects, &ms.textureIDs)

	// button callback function
	// starts from 1 so 0 is nil
	ms.callbacks = append(ms.callbacks, nil)
	ms.callbacks = append(ms.callbacks, menuToPlay)
	ms.callbacks = append(ms.callbacks, exitToMenu)

	ms.SetCallBacks(ms.callbacks)

	return true
}

// OnExit ...
func (ms *MainMenuState) OnExit() bool {
	// var err error

	fmt.Println("exit menu state")

	for _, v := range ms.textureIDs {
		STextureManager.ClearFromTextureMap(v)
	}

	return true
}

// GetStateID ...
func (ms MainMenuState) GetStateID() string {
	return MenuID
}

// SetCallBacks ...
func (ms *MainMenuState) SetCallBacks(cb []Callback) {
	// are of type IGameObject
	for _, v := range ms.objects {
		switch v.(type) {
		// if type menubutton
		case *MenuButton:
			fmt.Println("menu button")
			button := v.(*MenuButton)

			// set callback based on button loaded id
			// 1 - menuToPlay 2 - exitToMenu
			cb := ms.callbacks[button.GetCallBackID()]
			button.SetCallBack(cb)
		}
	}
}

// --- Callbacks
func menuToPlay() {
	fmt.Println("PLAY BUTTON CLICKED")
	STheGame.GetStateMachine().ChangeState(NewPlayState())
}

func exitToMenu() {
	fmt.Println("EXIT BUTTON CLICKED")
	STheGame.Quit()
}
