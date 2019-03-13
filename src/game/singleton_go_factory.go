package game

// GAME OBJECT FACTORY
// SINGLETON

import (
	"errors"
	"fmt"
	"sync"

	"../gologger"
)

// GOFactory ...
type GOFactory struct {
	GoCreator map[string]ICreator
}

var (
	gofactory *GOFactory
	fOnce     sync.Once
)

// STheGameObjFactory ...
var STheGameObjFactory = newGoFactory()

// NewGoFactory ... convert into a singleton
func newGoFactory() *GOFactory {
	gologger.SLogger.Println("Init Game Object Factory Singleton")
	fOnce.Do(func() {
		gofactory = &GOFactory{
			GoCreator: make(map[string]ICreator),
		}
	})
	return gofactory
}

// Register ...
func (gf *GOFactory) Register(typeID string, creator ICreator) bool {
	fmt.Println("registering", typeID)

	// check if already registered
	_, ok := gf.GoCreator[typeID]
	if ok {
		fmt.Println("already registered obj")

		return false
	}

	gf.GoCreator[typeID] = creator

	gologger.SLogger.Println("Added To Factory Obj Of Type", typeID)

	return true
}

// Create ...
func (gf *GOFactory) Create(typeID string) (IGameObject, error) {
	v, ok := gf.GoCreator[typeID]

	// not found
	if !ok {
		return nil, errors.New("factory object not found")
	}

	bc := v.(ICreator)

	gologger.SLogger.Println("Factory Created Obj Of Type", typeID)

	// call its create function
	return bc.CreateObj(), nil
}
