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

// PauseState ...
type PauseState struct {
	pauseID     string
	gameObjects []IGameObject
}

// NewPauseState ...
func NewPauseState() *PauseState {
	ps := &PauseState{}
	ps.pauseID = "PAUSE"
	return ps
}

// Update ...
func (ps *PauseState) Update() {
	for _, v := range ps.gameObjects {
		v.Update()
	}
}

// Render ...
func (ps *PauseState) Render() {
	for _, v := range ps.gameObjects {
		v.Draw()
	}
}

// OnEnter ...
func (ps *PauseState) OnEnter() bool {
	fmt.Println("enter pause state")

	// load textures
	STextureManager.Load("assets/main.png", "mainbutton", STheGame.GetRenderer())
	STextureManager.Load("assets/resume.png", "resumebutton", STheGame.GetRenderer())

	// buttons
	// main - go back to main state
	// resume - go back to previous state
	mainbutton := NewMenuButton(NewParams(200, 300, 200, 100, "mainbutton", 0), func() {
		fmt.Println("MENU BUTTON CLICKED")
		// go to menu state if clicked
		STheGame.GetStateMachine().ChangeState(NewMenuState())
	})

	resumebutton := NewMenuButton(NewParams(200, 100, 200, 100, "resumebutton", 0), func() {
		fmt.Println("RESUME BUTTON CLICKED")
		// pop this state and go to previous
		STheGame.GetStateMachine().PopState()
	})

	// add to gameobjects
	ps.gameObjects = append(ps.gameObjects, resumebutton)
	ps.gameObjects = append(ps.gameObjects, mainbutton)

	return true
}

// OnExit ...
func (ps *PauseState) OnExit() bool {
	fmt.Println("exit pause state")

	for _, v := range ps.gameObjects {
		v.Clean()
	}

	ps.gameObjects = nil

	STextureManager.ClearFromTextureMap("resumebutton")
	STextureManager.ClearFromTextureMap("mainbutton")

	SInputHandler.Reset()

	return true
}

// GetStateID ...
func (ps PauseState) GetStateID() string {
	return ps.pauseID
}
