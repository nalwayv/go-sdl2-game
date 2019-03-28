package game

/*
* Info
* ---
* Used to create blank object of Animated Graphic
*
* ICreator interface
* ---
*
* - CreateObj() IGameObject
**/

// AnimGraphicCreator ...
type AnimGraphicCreator struct{}

// NewAnimGraphicCreator ...
func NewAnimGraphicCreator() *AnimGraphicCreator {
	return &AnimGraphicCreator{}
}

// CreateObj ...
func (ac *AnimGraphicCreator) CreateObj() IGameObject {
	return NewAnimatedGraphic()
}
