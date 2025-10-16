
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [itemIndex] class.
var itemIndexClass _itemIndexClass

func init() {
	itemIndexClass = _itemIndexClass{objc.GetClass("itemIndex")}
}

type _itemIndexClass struct {
	objc.Class
}

// An interface definition for the [itemIndex] class.
type IitemIndex interface {
	ID() objc.ID
}

type itemIndex struct {
	id objc.ID
}

func itemIndexFrom(ptr unsafe.Pointer) itemIndex {
	return itemIndex{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ itemIndex) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _itemIndexClass) Alloc() itemIndex {
	rv := objc.Send[itemIndex](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _itemIndexClass) New() itemIndex {
	rv := objc.Send[itemIndex](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewitemIndex creates and returns a new initialized instance.
func NewitemIndex() itemIndex {
	return itemIndexClass.New()
}

// Init initializes the instance.
func (i_ itemIndex) Init() itemIndex {
	rv := objc.Send[itemIndex](i_.ID(), selInit)
	return rv
}
