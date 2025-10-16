
// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [isSelectable] class.
var isSelectableClass _isSelectableClass

func init() {
	isSelectableClass = _isSelectableClass{objc.GetClass("isSelectable")}
}

type _isSelectableClass struct {
	objc.Class
}

// An interface definition for the [isSelectable] class.
type IisSelectable interface {
	ID() objc.ID
}

type isSelectable struct {
	id objc.ID
}

func isSelectableFrom(ptr unsafe.Pointer) isSelectable {
	return isSelectable{
		id: objc.ID(ptr),
	}
}

// ID returns the underlying objc.ID.
func (i_ isSelectable) ID() objc.ID {
	return i_.id
}

// Alloc allocates a new instance without initialization.
func (ic _isSelectableClass) Alloc() isSelectable {
	rv := objc.Send[isSelectable](objc.ID(ic.Class), selAlloc)
	return rv
}

// New creates and returns a new initialized instance.
func (ic _isSelectableClass) New() isSelectable {
	rv := objc.Send[isSelectable](objc.ID(ic.Class), selNew)
	objc.Send[objc.ID](rv.ID(), selAutorelease)
	return rv
}

// NewisSelectable creates and returns a new initialized instance.
func NewisSelectable() isSelectable {
	return isSelectableClass.New()
}

// Init initializes the instance.
func (i_ isSelectable) Init() isSelectable {
	rv := objc.Send[isSelectable](i_.ID(), selInit)
	return rv
}
