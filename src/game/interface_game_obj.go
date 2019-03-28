package game

/*
* Info
* ---
* Interface for game object classes
**/

// IGameObject ...
type IGameObject interface {
	Draw()
	Update()
	Clean()
	Load(*LoadParams)
}
