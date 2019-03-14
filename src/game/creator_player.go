package game

/*
ICreator interface
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
