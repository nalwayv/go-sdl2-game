package game

/*
* Used on classes that create a blank game object
* for the game object factory
**/

// ICreator ...
type ICreator interface {
	CreateObj() IGameObject
}
