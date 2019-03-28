package game

/*
* Info
* ---
* Used to create blank object of Enemy
*
* ICreator interface
* ---
*
* - CreateObj() IGameObject
**/

// EnemyCreator ...
type EnemyCreator struct{}

// NewEnemyCreator ...
func NewEnemyCreator() *EnemyCreator {
	return &EnemyCreator{}
}

// CreateObj ...
func (pc *EnemyCreator) CreateObj() IGameObject {
	return NewEnemy()
}
