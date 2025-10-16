
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [itemAlignment] class.
var itemAlignmentClass _itemAlignmentClass

func init() {
	itemAlignmentClass = _itemAlignmentClass{objc.GetClass("itemAlignment")}
}

type _itemAlignmentClass struct {
	objc.Class
}

// An interface definition for the [itemAlignment] class.
type IitemAlignment interface {
	ID() objc.ID
}

type itemAlignment struct {
	id objc.ID
}

func itemAlignmentFrom(ptr unsafe.Pointer) itemAlignment {
	return itemAlignment{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ itemAlignment) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _itemAlignmentClass) Alloc() itemAlignment {
	rv := objc.Send[itemAlignment](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _itemAlignmentClass) New() itemAlignment {
	rv := objc.Send[itemAlignment](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewitemAlignment creates and returns a new initialized instance.
func NewitemAlignment() itemAlignment {
	return itemAlignmentClass.New()
}

// Init initializes the instance.
func (i_ itemAlignment) Init() itemAlignment {
	rv := objc.Send[itemAlignment](i_.ID(), selInit)
	return rv
}
