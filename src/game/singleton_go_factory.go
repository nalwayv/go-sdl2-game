package game


/*
* Singleton / Object Factory
* ---
* Object factory that crates blank objects that get filled out
* then a stats onEnter is called passing in info parsed from data file
**/

import (
	"errors"
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

// Register ... add a new blank object to the map that implements ICreator interface
func (gf *GOFactory) Register(typeID string, creator ICreator) bool {
	gologger.SLogger.Println("registering", typeID)

	// check if already registered
	_, ok := gf.GoCreator[typeID]
	if ok {
		gologger.SLogger.Println("Already Registered Object ", typeID)

		return false
	}

	gf.GoCreator[typeID] = creator

	gologger.SLogger.Println("Added To Factory Obj Of Type", typeID)

	return true
}

// Create ... find object within map and call its createOBJ else raise error
func (gf *GOFactory) Create(typeID string) (IGameObject, error) {
	v, ok := gf.GoCreator[typeID]

	// not found
	if !ok {
		return nil, errors.New("factory object not found " + typeID)
	}

	bc := v.(ICreator)

	gologger.SLogger.Println("Factory Created Obj Of Type", typeID)

	// call its create function
	return bc.CreateObj(), nil
}
