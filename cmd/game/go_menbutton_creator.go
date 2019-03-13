package game

/*
Implements ICreator interface

- CreateObj() IGameObject
*/

// MenuButtonCreator ...
type MenuButtonCreator struct{}

// NewMenuButtonCreator ...
func NewMenuButtonCreator() *MenuButtonCreator {
	return &MenuButtonCreator{}
}

// CreateObj ...
func (mb *MenuButtonCreator) CreateObj() IGameObject {
	return NewMenuButton()
}
