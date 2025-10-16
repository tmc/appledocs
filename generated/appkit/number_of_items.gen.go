
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [numberOfItems] class.
var numberOfItemsClass _numberOfItemsClass

func init() {
	numberOfItemsClass = _numberOfItemsClass{objc.GetClass("numberOfItems")}
}

type _numberOfItemsClass struct {
	objc.Class
}

// An interface definition for the [numberOfItems] class.
type InumberOfItems interface {
	ID() objc.ID
}

type numberOfItems struct {
	id objc.ID
}

func numberOfItemsFrom(ptr unsafe.Pointer) numberOfItems {
	return numberOfItems{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (n_ numberOfItems) ID() objc.ID {
	return n_.id
}

// Alloc allocates a new instance without initialization.
func (nc _numberOfItemsClass) Alloc() numberOfItems {
	rv := objc.Send[numberOfItems](objc.ID(nc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (nc _numberOfItemsClass) New() numberOfItems {
	rv := objc.Send[numberOfItems](objc.ID(nc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewnumberOfItems creates and returns a new initialized instance.
func NewnumberOfItems() numberOfItems {
	return numberOfItemsClass.New()
}

// Init initializes the instance.
func (n_ numberOfItems) Init() numberOfItems {
	rv := objc.Send[numberOfItems](n_.ID(), selInit)
	return rv
}
