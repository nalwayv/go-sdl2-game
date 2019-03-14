package game


/*
ICreator interface
---

- CreateObj() IGameObject
*/

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
