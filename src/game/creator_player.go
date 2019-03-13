package game

/*
Implements ICreator interface

- CreateObj() IGameObject
*/

// PlayerCreator ...
type PlayerCreator struct{}

// NewPlayerCreator ...
func NewPlayerCreator() *PlayerCreator {
	return &PlayerCreator{}
}

// CreateObject ...
func (pc *PlayerCreator) CreateObject() IGameObject {
	player := Player{}
	return &player
}
