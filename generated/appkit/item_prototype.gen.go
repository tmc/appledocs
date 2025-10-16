
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [itemPrototype] class.
var itemPrototypeClass _itemPrototypeClass

func init() {
	itemPrototypeClass = _itemPrototypeClass{objc.GetClass("itemPrototype")}
}

type _itemPrototypeClass struct {
	objc.Class
}

// An interface definition for the [itemPrototype] class.
type IitemPrototype interface {
	ID() objc.ID
}

type itemPrototype struct {
	id objc.ID
}

func itemPrototypeFrom(ptr unsafe.Pointer) itemPrototype {
	return itemPrototype{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ itemPrototype) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _itemPrototypeClass) Alloc() itemPrototype {
	rv := objc.Send[itemPrototype](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _itemPrototypeClass) New() itemPrototype {
	rv := objc.Send[itemPrototype](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewitemPrototype creates and returns a new initialized instance.
func NewitemPrototype() itemPrototype {
	return itemPrototypeClass.New()
}

// Init initializes the instance.
func (i_ itemPrototype) Init() itemPrototype {
	rv := objc.Send[itemPrototype](i_.ID(), selInit)
	return rv
}
