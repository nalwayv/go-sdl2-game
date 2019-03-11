package game

/*
Implements ICreator interface

- CreateObj() IGameObject
*/

// CreatePlayer ...
type CreatePlayer struct {
}

// CreateObject ...
func (c *CreatePlayer) CreateObject() IGameObject {
	player := Player{}
	return &player
}
