
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [itemIdentifiers] class.
var itemIdentifiersClass _itemIdentifiersClass

func init() {
	itemIdentifiersClass = _itemIdentifiersClass{objc.GetClass("itemIdentifiers")}
}

type _itemIdentifiersClass struct {
	objc.Class
}

// An interface definition for the [itemIdentifiers] class.
type IitemIdentifiers interface {
	ID() objc.ID
}

type itemIdentifiers struct {
	id objc.ID
}

func itemIdentifiersFrom(ptr unsafe.Pointer) itemIdentifiers {
	return itemIdentifiers{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ itemIdentifiers) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _itemIdentifiersClass) Alloc() itemIdentifiers {
	rv := objc.Send[itemIdentifiers](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _itemIdentifiersClass) New() itemIdentifiers {
	rv := objc.Send[itemIdentifiers](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewitemIdentifiers creates and returns a new initialized instance.
func NewitemIdentifiers() itemIdentifiers {
	return itemIdentifiersClass.New()
}

// Init initializes the instance.
func (i_ itemIdentifiers) Init() itemIdentifiers {
	rv := objc.Send[itemIdentifiers](i_.ID(), selInit)
	return rv
}
