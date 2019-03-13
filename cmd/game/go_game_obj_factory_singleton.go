package game

// GAME OBJECT FACTORY
// SINGLETON

import (
	"errors"
	"fmt"
	"sync"
)

var (
	factory *GOFactory
	fOnce   sync.Once
)

// STheGameObjFactory ...
var STheGameObjFactory = newGoFactory()

// GOFactory ...
type GOFactory struct {
	GoCreator map[string]ICreator
}

// NewGoFactory ...
func newGoFactory() *GOFactory {

	fOnce.Do(func() {
		factory := &GOFactory{}
		factory.GoCreator = make(map[string]ICreator)
	})
	return factory
}

// Register ...
func (gf *GOFactory) Register(typeID string, creator ICreator) bool {
	// check if already registered
	// else add
	fmt.Println("registering obj")
	_, ok := gf.GoCreator[typeID]
	if ok {
		fmt.Println("already registered obj")

		return false
	}

	fmt.Println("Registered obj of type", typeID)
	gf.GoCreator[typeID] = creator
	return true
}

// Create ...
func (gf *GOFactory) Create(typeID string) (IGameObject, error) {
	v, ok := gf.GoCreator[typeID]

	// not found
	if !ok {
		return nil, errors.New("factory object not found")
	}

	fmt.Println("Created obj of type", typeID)

	return v.CreateObj(), nil
}
