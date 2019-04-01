package game

/*
Info
---
Used to create blank object of Menu Button

ICreator interface
---
- CreateObj() IGameObject
**/

// MenuButtonCreator ...
type MenuButtonCreator struct{}

// NewMenuButtonCreator ...
func NewMenuButtonCreator() *MenuButtonCreator {
	return &MenuButtonCreator{}
}

// CreateObj ...
func (mc *MenuButtonCreator) CreateObj() IGameObject {
	return NewMenuButton()
}
