package game

/*
Info
---
Used to create an blank base object that will be placed into the
game object factory


Implements ICreator interface
---
- CreateObj() IGameObject
*/

// PlayerCreator ...
type PlayerCreator struct{}

// NewPlayerCreator ...
func NewPlayerCreator() *PlayerCreator {
	return &PlayerCreator{}
}

// CreateObj ...
func (pc *PlayerCreator) CreateObj() IGameObject {
	return NewPlayer()
}
