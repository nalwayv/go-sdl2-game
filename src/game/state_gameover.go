package game

/*
Implements IGameState interface.

- Update()
- Render()
- OnEnter() bool
- OnExit() bool
- GetStateID() string
*/
import "fmt"

// GameOver ...
type GameOver struct {
	gameObjects []IGameObject

	playID string
}

// NewGameOverState ...
func NewGameOverState() *GameOver {
	ps := &GameOver{}
	ps.gameObjects = make([]IGameObject, 0)
	ps.playID = "GAMEOVER"
	return ps
}

// Update ...
func (ps *GameOver) Update() {
	for _, v := range ps.gameObjects {
		v.Update()
	}
}

// Render ...
func (ps *GameOver) Render() {
	for _, v := range ps.gameObjects {
		v.Draw()
	}
}

// OnEnter ...
func (ps *GameOver) OnEnter() bool {
	fmt.Println("enter gameover state")

	// load textures
	STextureManager.Load("assets/gameover.png", "gameovertext", STheGame.GetRenderer())
	STextureManager.Load("assets/main.png", "mainbutton", STheGame.GetRenderer())
	STextureManager.Load("assets/restart.png", "restartbutton", STheGame.GetRenderer())

	// set buttons / functions
	// mainbutton := NewMenuButton(NewParams(200, 200, 200, 80, "mainbutton", 0, 0, 0), func() {
	// 	STheGame.GetStateMachine().ChangeState(NewMenuState())
	// })

	// restartbutton := NewMenuButton(NewParams(200, 300, 200, 80, "restartbutton", 0, 0, 0), func() {
	// 	STheGame.GetStateMachine().ChangeState(NewPlayState())
	// })

	// set text
	gameovertext := NewAnimatedGraphic()

	// add to gameobjects slice
	// ps.gameObjects = append(ps.gameObjects, mainbutton)
	// ps.gameObjects = append(ps.gameObjects, restartbutton)
	ps.gameObjects = append(ps.gameObjects, gameovertext)

	return true
}

// OnExit ...
func (ps *GameOver) OnExit() bool {
	var err error

	fmt.Println("exit gameover state")

	for _, v := range ps.gameObjects {
		v.Clean()
	}

	ps.gameObjects = nil

	err = STextureManager.ClearFromTextureMap("mainbutton")
	checkError(err)
	err = STextureManager.ClearFromTextureMap("restartbutton")
	checkError(err)
	err = STextureManager.ClearFromTextureMap("gameovertext")
	checkError(err)

	return true
}

// GetStateID ... get state id
func (ps GameOver) GetStateID() string {
	return ps.playID
}
