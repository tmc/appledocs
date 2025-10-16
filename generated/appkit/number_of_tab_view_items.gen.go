
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [numberOfTabViewItems] class.
var numberOfTabViewItemsClass _numberOfTabViewItemsClass

func init() {
	numberOfTabViewItemsClass = _numberOfTabViewItemsClass{objc.GetClass("numberOfTabViewItems")}
}

type _numberOfTabViewItemsClass struct {
	objc.Class
}

// An interface definition for the [numberOfTabViewItems] class.
type InumberOfTabViewItems interface {
	ID() objc.ID
}

type numberOfTabViewItems struct {
	id objc.ID
}

func numberOfTabViewItemsFrom(ptr unsafe.Pointer) numberOfTabViewItems {
	return numberOfTabViewItems{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (n_ numberOfTabViewItems) ID() objc.ID {
	return n_.id
}

// Alloc allocates a new instance without initialization.
func (nc _numberOfTabViewItemsClass) Alloc() numberOfTabViewItems {
	rv := objc.Send[numberOfTabViewItems](objc.ID(nc.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (nc _numberOfTabViewItemsClass) New() numberOfTabViewItems {
	rv := objc.Send[numberOfTabViewItems](objc.ID(nc.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewnumberOfTabViewItems creates and returns a new initialized instance.
func NewnumberOfTabViewItems() numberOfTabViewItems {
	return numberOfTabViewItemsClass.New()
}

// Init initializes the instance.
func (n_ numberOfTabViewItems) Init() numberOfTabViewItems {
	rv := objc.Send[numberOfTabViewItems](n_.ID(), selInit)
	return rv
}
