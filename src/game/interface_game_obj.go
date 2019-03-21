package game

/*
Info
---
Interface for a game object
**/

// IGameObject ...
type IGameObject interface {
	Draw()
	Update()
	Clean()
	Load(*LoadParams)
}
